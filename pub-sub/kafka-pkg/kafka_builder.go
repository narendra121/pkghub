package kafkapkg

type KafkaCfg struct {
	TopicInfo     TopicCfgBuilder
	ProducerInfo  ProducerCfgBuilder
	ConsumerGroup ConsumerGroupCfgBuilder
}

type KafkaBuilder struct {
	kafkacfg KafkaCfg
}

func NewKafkaCfgBuilder() *KafkaBuilder {
	return &KafkaBuilder{kafkacfg: KafkaCfg{TopicInfo: TopicCfgBuilder{}, ProducerInfo: ProducerCfgBuilder{}, ConsumerGroup: ConsumerGroupCfgBuilder{}}}
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

func (kb *KafkaBuilder) Build() KafkaCfg {
	return kb.kafkacfg
}
