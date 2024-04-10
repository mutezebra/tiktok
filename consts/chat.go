package consts

// msg_type
const (
	ChatTextMessage    = 1
	HistoryChatMessage = 2
	NotReadMessage     = 3

	ChatMQReadChSize              = 100
	ChatMQWriteChSize             = 100
	ChatMQPersiTopicName          = "chat-msg-topic"
	ChatMQPersiTopicPartitions    = 3
	ChatMQPersiReplicationFactors = 1

	ChatMQPersiReaderGroupName   = "chat-msg-reader-group"
	ChatMQPersiReaderGroupNumber = 3

	ChatDefaultPersistenceGoroutineNum = 10
)
