package chat

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/app/interface/persistence/database"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"sync"
)

type Service struct {
	Manager *Manager
	repo    repository.ChatRepository
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
		}
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

func (s *Service) Send(msg *Message, from, to int64) error {
	var err error
	switch msg.MsgType {
	case consts.ChatTextMessage:
		err = s.Manager.Send([]byte(msg.Content), to)
		break
	default:
		return s.WriteError("not supported msg type", from)
	}
	return err
}

func (s *Service) WriteError(msg string, to int64) error {
	return s.Manager.Send([]byte(msg), to)
}
