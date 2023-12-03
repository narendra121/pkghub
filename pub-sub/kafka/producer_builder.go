package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

/*
-------------------------------------------Usage----------------------

	pb := kafka.NewProducerBuilder()
	pb.SetBrokers([]string{"kafka:9092"}).
		SetAckType(sarama.WaitForAll).
		Build()

	kb := kafka.NewKafkaBuilder()

	kb.SetProducerInfo(*pb).Build()
kf := kafka.NewPubSubFactory(kb)

producer, err := kf.CreateSyncProducer()
	log.Println(err)
	kf.SendMessage(producer, "topic","helwlo")


*/

type Producer struct {
	Brokers []string
	Config  *sarama.Config
}

type ProducerBuilder struct {
	producerCfg Producer
}

func NewProducerBuilder() *ProducerBuilder {
	return &ProducerBuilder{producerCfg: Producer{Config: sarama.NewConfig()}}
}

func (pb *ProducerBuilder) SetAckType(ackType sarama.RequiredAcks) *ProducerBuilder {
	pb.producerCfg.Config.Producer.RequiredAcks = ackType
	return pb
}

func (pb *ProducerBuilder) SetCompressionType(compressionType sarama.CompressionCodec) *ProducerBuilder {
	pb.producerCfg.Config.Producer.Compression = compressionType
	return pb
}

func (pb *ProducerBuilder) SetKafkaBackoffRetry(retryCount, maxRetryCount, backOffMultiplier int64, initBackoff, maxBackOff time.Duration) *ProducerBuilder {
	pb.producerCfg.Config.Producer.Retry.BackoffFunc = func(retries, maxRetries int) time.Duration {
		if retries == 0 {
			return initBackoff
		}
		for i := 0; i < maxRetries; i++ {
			initBackoff = initBackoff * time.Duration(backOffMultiplier)
			if initBackoff > maxBackOff {
				return maxBackOff
			}
		}
		return initBackoff
	}
	return pb
}

func (pb *ProducerBuilder) SetBrokers(brokers []string) *ProducerBuilder {
	pb.producerCfg.Brokers = append(pb.producerCfg.Brokers, brokers...)
	return pb
}

func (pb *ProducerBuilder) Build() Producer {
	pb.producerCfg.Config.Producer.Return.Successes = true
	pb.producerCfg.Config.Producer.Return.Errors = true
	return pb.producerCfg
}
