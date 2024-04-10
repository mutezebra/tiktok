package chat

import (
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/pkg/errors"
)

const (
	topicName          = consts.ChatMQPersiTopicName
	partitions         = consts.ChatMQPersiTopicPartitions
	replicationFactors = consts.ChatMQPersiReplicationFactors
	readerGroup        = consts.ChatMQPersiReaderGroupName
	readerNum          = consts.ChatMQPersiReaderGroupNumber
)

func (s *Service) EnableMQ() {
	if err := s.mq.CreateTopic(topicName, partitions, replicationFactors); err != nil {
		log.LogrusObj.Panic(errors.Wrap(err, "create topic failed"))
	}
	go s.mq.RunWriter(s.ctx, topicName, s.mqWriteCh)
	go s.mq.RunGroupReader(s.ctx, topicName, readerGroup, readerNum, s.mqReadCh)
}

func (s *Service) writeMsgToMQ(msg *Message) error {
	data, err := s.getMsgBytes(msg)
	if err != nil {
		return err
	}
	s.mqWriteCh <- data
	return nil
}

func (s *Service) readMsgFromMQ() (*Message, error) {
	data := <-s.mqReadCh
	return s.getMsgFromBytes(data)
}
