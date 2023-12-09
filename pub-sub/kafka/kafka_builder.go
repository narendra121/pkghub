package kafka

type Kafka struct {
	TopicInfo     TopicCfgBuilder
	ProducerInfo  ProducerCfgBuilder
	ConsumerGroup ConsumerGroupCfgBuilder
}

type KafkaBuilder struct {
	kafkacfg Kafka
}

func NewKafkaCfgBuilder() *KafkaBuilder {
	return &KafkaBuilder{kafkacfg: Kafka{TopicInfo: TopicCfgBuilder{}, ProducerInfo: ProducerCfgBuilder{}, ConsumerGroup: ConsumerGroupCfgBuilder{}}}
}

func (kb *KafkaBuilder) SetTopicInfo(topicInfo TopicCfgBuilder) *KafkaBuilder {
	kb.kafkacfg.TopicInfo = topicInfo
	return kb
}

func (kb *KafkaBuilder) SetProducerInfo(producerInfo ProducerCfgBuilder) *KafkaBuilder {
	kb.kafkacfg.ProducerInfo = producerInfo
	return kb
}

func (kb *KafkaBuilder) SetConsumerGroupInfo(consumerGroupInfo ConsumerGroupCfgBuilder) *KafkaBuilder {
	kb.kafkacfg.ConsumerGroup = consumerGroupInfo
	return kb
}

func (kb *KafkaBuilder) Build() Kafka {
	return kb.kafkacfg
}
