package kafka

import (
	"context"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

func (k *KafkaBuilder) CreateNewTopic() error {
	config := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(k.kafkaSetUp.TopicInfo.topic.Brokers, config)
	if err != nil {
		log.Errorln("error in CreateTopic ", err)
		return err
	}
	defer admin.Close()
	return admin.CreateTopic(k.kafkaSetUp.TopicInfo.topic.TopicName, k.kafkaSetUp.TopicInfo.topic.TopicConfig, false)
}

func (k *KafkaBuilder) CreateSyncProducer() (interface{}, error) {
	producer, err := sarama.NewSyncProducer(k.kafkaSetUp.ProducerInfo.producerCfg.Brokers, k.kafkaSetUp.ProducerInfo.producerCfg.Config)
	if err != nil {
		log.Errorln("error in producer ", err)
		return nil, err
	}

	return producer, nil
}

func (k *KafkaBuilder) SendMessage(producer interface{}, topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	_, _, err := producer.(sarama.SyncProducer).SendMessage(msg)
	if err != nil {
		log.Errorln("Failed to send message: ", err)
	}
}

func (k *KafkaBuilder) CreateConsumerGroup() (interface{}, error) {
	consumer, err := sarama.NewConsumerGroup(k.kafkaSetUp.ConsumerGroup.consumerCfg.Brokers, k.kafkaSetUp.ConsumerGroup.consumerCfg.GroupId, k.kafkaSetUp.ConsumerGroup.consumerCfg.Config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (k *KafkaBuilder) AddConsumerToConsumerGroup(consumer interface{}, customMsgHandler interface{}) {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	ctx := context.Background()

	go func() {
		for {
			if err := consumer.(sarama.ConsumerGroup).Consume(ctx, k.kafkaSetUp.ConsumerGroup.consumerCfg.Topics, customMsgHandler.(*MessageHandler)); err != nil {
				log.Errorln("Error from consumer:", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-signals
}
