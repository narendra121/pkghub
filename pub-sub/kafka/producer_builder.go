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

Producer, err := kf.CreateSyncProducer()
	log.Println(err)
	kf.SendMessage(Producer, "topic","helwlo")


*/

type ProducerCfg struct {
	Brokers []string
	Config  *sarama.Config
}

type ProducerCfgBuilder struct {
	producerCfg ProducerCfg
}

func NewProducerCfgBuilder() *ProducerCfgBuilder {
	return &ProducerCfgBuilder{producerCfg: ProducerCfg{Config: sarama.NewConfig()}}
}

func (pb *ProducerCfgBuilder) SetAckType(ackType sarama.RequiredAcks) *ProducerCfgBuilder {
	pb.producerCfg.Config.Producer.RequiredAcks = ackType
	return pb
}

func (pb *ProducerCfgBuilder) SetCompressionType(compressionType sarama.CompressionCodec) *ProducerCfgBuilder {
	pb.producerCfg.Config.Producer.Compression = compressionType
	return pb
}

func (pb *ProducerCfgBuilder) SetKafkaBackoffRetry(retryCount, maxRetryCount, backOffMultiplier int64, initBackoff, maxBackOff time.Duration) *ProducerCfgBuilder {
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

func (pb *ProducerCfgBuilder) SetBrokers(brokers []string) *ProducerCfgBuilder {
	pb.producerCfg.Brokers = append(pb.producerCfg.Brokers, brokers...)
	return pb
}

func (pb *ProducerCfgBuilder) Build() ProducerCfg {
	pb.producerCfg.Config.Producer.Return.Successes = true
	pb.producerCfg.Config.Producer.Return.Errors = true
	return pb.producerCfg
}
