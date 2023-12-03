package kafka

import "github.com/IBM/sarama"

type customMessageHandlerFunc func(*sarama.ConsumerMessage)
type MessageHandler struct {
	CustomMessageHandler customMessageHandlerFunc
}

func (mh *MessageHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (mh *MessageHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (mh *MessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		mh.CustomMessageHandler(message)
		session.MarkMessage(message, "")
	}
	return nil
}
