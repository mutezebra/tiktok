package kafka

import (
	"context"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"io"
	"os"
	"os/signal"
)

type MQModel struct {
}

func (m *MQModel) CreateTopic(topic string, partitions int, replicationFactors int) error {
	conn, err := kafka.Dial(config.Conf.Kafka.Network, config.Conf.Kafka.Address)
	if err != nil {
		return errors.Wrap(err, "failed to dial kafka")
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.LogrusObj.Error(errors.Wrap(err, "failed to close kafka conn"))
		}
	}()

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactors,
	})
	if err != nil {
		return errors.Wrap(err, "create topic failed")
	}

	return nil
}

func (m *MQModel) RunGroupReader(ctx context.Context, topic string, groupID string, readerNumber int, ch chan []byte) {
	defer close(ch)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	closeFns := make([]func(), 0, readerNumber)
	defer closeAllConn(closeFns)

	cfg := kafka.ReaderConfig{
		Brokers:     []string{config.Conf.Kafka.Address},
		GroupID:     groupID,
		Topic:       topic,
		MinBytes:    consts.ReaderDefaultMinBytes,
		MaxBytes:    consts.ReaderDefaultMaxBytes,
		MaxAttempts: consts.ReaderDefaultAttempts,
	}

	for i := 0; i < readerNumber; i++ {
		go func() {
			r := kafka.NewReader(cfg)
			closeFns = append(closeFns, closeConn(r))
			for {
				msg, err := r.ReadMessage(ctx)
				if err != nil {
					log.LogrusObj.Error(errors.Wrap(err, "failed read message from kafka"))
					break
				}
				ch <- msg.Value
			}
		}()
	}

	<-interrupt
	return
}

func (m *MQModel) RunWriter(ctx context.Context, topic string, ch chan []byte) {
	defer close(ch)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	closeFns := make([]func(), 0, 1)
	defer closeAllConn(closeFns)

	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Conf.Kafka.Address),
		Topic:                  topic,
		Balancer:               &kafka.RoundRobin{},
		MaxAttempts:            consts.WriterDefaultAttempts,
		RequiredAcks:           kafka.RequireOne,
		Async:                  true,
		AllowAutoTopicCreation: false,
	}
	closeFns = append(closeFns, closeConn(w))

	go func() {
		for {
			msg, ok := <-ch
			if !ok {
				break
			}
			err := w.WriteMessages(ctx, kafka.Message{Value: msg})
			if err != nil && err != io.EOF {
				log.LogrusObj.Error(errors.Wrap(err, "failed write message to kafka"))
				break
			}
		}
	}()

	<-interrupt
	return
}
