package chat

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/Mutezebra/tiktok/app/domain/model"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/app/interface/persistence/database"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/kafka"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"sync"
)

type Service struct {
	Manager *Manager
	repo    repository.ChatRepository

	mq        model.MQ
	mqReadCh  chan []byte
	mqWriteCh chan []byte

	ctx        context.Context
	bufferPool sync.Pool
}

var once sync.Once
var defaultService *Service

func DefaultService() *Service {
	once.Do(func() {
		defaultService = &Service{
			Manager: &Manager{
				connMap: make(map[int64]*websocket.Conn),
				msgChan: make(map[int64]chan []byte),
				repo:    database.NewChatRepository(),
			},
			repo: database.NewChatRepository(),

			mq:        &kafka.MQModel{},
			mqReadCh:  make(chan []byte, consts.ChatMQReadChSize),
			mqWriteCh: make(chan []byte, consts.ChatMQWriteChSize),

			ctx:        context.Background(),
			bufferPool: sync.Pool{New: func() any { return new(bytes.Buffer) }},
		}

		defaultService.EnableSyncPersistence()
	})
	return defaultService
}

func (s *Service) NewClient(from, to int64, conn *websocket.Conn) *Client {
	return &Client{
		srv:     s,
		From:    from,
		To:      to,
		Channel: make(chan []byte),
		Conn:    conn,
	}
}

func (s *Service) Register(ctx context.Context, client *Client) error {
	exist, err := s.repo.WhetherExistUser(ctx, client.To)
	if !exist {
		return errors.WithMessage(err, "user not exist")
	}
	if err != nil {
		return err
	}

	s.Manager.Register(client.From, client.Conn, client.Channel)
	return nil
}

func (s *Service) Unregister(client *Client) error {
	err := client.Conn.Close()
	if err != nil {
		return errors.New("websocket conn close failed")
	}

	close(client.Channel)
	s.Manager.Unregister(client.From)
	return nil
}

func (s *Service) SendMsg(msg *Message) error {
	var err error
	switch msg.MsgType {
	case consts.ChatTextMessage:
		if err = s.Manager.Send([]byte(msg.Content), msg.To); err != nil {
			break
		}
		if err = s.writeMsgToMQ(msg); err != nil {
			break
		}
		break
	default:
		return s.WriteErrorToConn("not supported msg type", msg.From)
	}
	return err
}

func (s *Service) WriteErrorToConn(error string, to int64) error {
	return s.Manager.Send([]byte(error), to)
}

func (s *Service) getMsgBytes(msg *Message) ([]byte, error) {
	buf := s.getBuffer()
	defer s.putBuffer(buf)

	enc := gob.NewEncoder(buf)
	err := enc.Encode(msg)
	if err != nil {
		return nil, errors.Wrap(err, "gob encoder encode msg failed")
	}

	return buf.Bytes(), nil
}

func (s *Service) getMsgFromBytes(data []byte) (*Message, error) {
	buf := s.getBuffer()
	defer s.putBuffer(buf)

	if _, err := buf.Write(data); err != nil {
		return nil, errors.Wrap(err, "failed write data To buffer")
	}

	dec := gob.NewDecoder(buf)
	msg := &Message{}
	if err := dec.Decode(msg); err != nil {
		return nil, errors.Wrap(err, "gob decoder decode msg failed")
	}

	return msg, nil
}

func (s *Service) getBuffer() *bytes.Buffer {
	buf := s.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (s *Service) putBuffer(buf *bytes.Buffer) {
	s.bufferPool.Put(buf)
}

func (s *Service) EnableSyncPersistence() {
	s.EnableMQ()
	go s.MessagePersistence(consts.ChatDefaultPersistenceGoroutineNum)
}

func (s *Service) MessagePersistence(asyncNumber int) {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	var convert = func(message *Message) *repository.Message {
		return &repository.Message{
			UID:        message.From,
			ReceiverID: message.To,
			Content:    message.Content,
			CreateAt:   message.CreateAt,
			DeleteAt:   0,
		}
	}

	chs := make([]chan *repository.Message, 0, asyncNumber)
	for i := 0; i < asyncNumber; i++ {
		ch := make(chan *repository.Message)
		chs = append(chs, ch)
		go func() {
			go s.repo.CreateMessageWithChannel(s.ctx, ch)
			for {
				msg, err := s.readMsgFromMQ()
				fmt.Println(msg.From)
				if err != nil {
					log.LogrusObj.Error(errors.WithMessage(err, "message persistence failed when read msg From mq"))
					return
				}
				repoMsg := convert(msg)
				ch <- repoMsg
			}
		}()
	}

	<-interrupt
	for _, ch := range chs {
		close(ch)
	}
	return
}
