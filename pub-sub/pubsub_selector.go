package pubsub

type PubSubFactory interface {
	CreateNewTopic() error

	CreateSyncProducer() (interface{}, error)
	SendMessage(producer interface{}, topic, message string)

	CreateConsumerGroup() (interface{}, error)
	AddConsumerToConsumerGroup(consumer interface{}, customMsgHandler interface{})
}
