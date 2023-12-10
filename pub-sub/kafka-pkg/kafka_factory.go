package kafkapkg

import (
	"context"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

type KafkaFactory interface {
	CreateNewTopic() error
	CreateSyncProducer() (sarama.SyncProducer, error)
	SendMessage(producer sarama.SyncProducer, topic, message string)
	CreateConsumerGroup() (sarama.ConsumerGroup, error)
	AddConsumerToConsumerGroup(consumer sarama.ConsumerGroup, customMsgHandler *MessageHandler)
}

func NewKafkaFactory() KafkaFactory {
	return &KafkaCfg{TopicInfo: TopicCfgBuilder{topicCfg: TopicCfg{Brokers: make([]string, 0), TopicConfig: &sarama.TopicDetail{}}},
		ProducerInfo:  ProducerCfgBuilder{producerCfg: ProducerCfg{Brokers: make([]string, 0), Config: sarama.NewConfig()}},
		ConsumerGroup: ConsumerGroupCfgBuilder{ConsumerGroupCfg{Brokers: make([]string, 0), Config: sarama.NewConfig()}},
	}
}

func (k *KafkaCfg) CreateNewTopic() error {
	config := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(k.TopicInfo.topicCfg.Brokers, config)
	if err != nil {
		log.Errorln("error in CreateTopic ", err)
		return err
	}
	defer admin.Close()
	return admin.CreateTopic(k.TopicInfo.topicCfg.TopicName, k.TopicInfo.topicCfg.TopicConfig, false)
}

func (k *KafkaCfg) CreateSyncProducer() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(k.ProducerInfo.producerCfg.Brokers, k.ProducerInfo.producerCfg.Config)
	if err != nil {
		log.Errorln("error in producer ", err)
		return nil, err
	}

	return producer, nil
}

func (k *KafkaCfg) SendMessage(producer sarama.SyncProducer, topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	_, _, err := producer.SendMessage(msg)
	if err != nil {
		log.Errorln("Failed to send message: ", err)
	}
}

func (k *KafkaCfg) CreateConsumerGroup() (sarama.ConsumerGroup, error) {
	consumer, err := sarama.NewConsumerGroup(k.ConsumerGroup.consumerCfg.Brokers, k.ConsumerGroup.consumerCfg.GroupId, k.ConsumerGroup.consumerCfg.Config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (k *KafkaCfg) AddConsumerToConsumerGroup(consumer sarama.ConsumerGroup, customMsgHandler *MessageHandler) {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	ctx := context.Background()

	go func() {
		for {
			if err := consumer.Consume(ctx, k.ConsumerGroup.consumerCfg.Topics, customMsgHandler); err != nil {
				log.Errorln("Error from consumer:", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-signals
}
