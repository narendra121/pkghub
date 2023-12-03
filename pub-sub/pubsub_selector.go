package pubsub

import (
	"pkg-hub/pub-sub/kafka"
)

type PubSubFactory interface {
	CreateNewTopic() error

	CreateSyncProducer() (interface{}, error)
	SendMessage(producer interface{}, topic, message string)

	CreateConsumerGroup() (interface{}, error)
	AddConsumerToConsumerGroup(consumer interface{}, customMsgHandler interface{})
}

func NewPubSubFactory(pubSubStreamer interface{}) PubSubFactory {

	switch pubSubStreamer.(type) {
	case *kafka.KafkaBuilder:
		return pubSubStreamer.(*kafka.KafkaBuilder)
	default:
		return nil
	}
}
