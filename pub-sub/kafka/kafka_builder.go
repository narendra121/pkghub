package kafka

type Kafka struct {
	TopicInfo     TopicBuilder
	ProducerInfo  ProducerBuilder
	ConsumerGroup ConsumerGroupBuilder
}

type KafkaBuilder struct {
	kafkaSetUp Kafka
}

func NewKafkaBuilder() *KafkaBuilder {
	return &KafkaBuilder{kafkaSetUp: Kafka{TopicInfo: TopicBuilder{}, ProducerInfo: ProducerBuilder{}, ConsumerGroup: ConsumerGroupBuilder{}}}
}

func (kb *KafkaBuilder) SetTopicInfo(topicInfo TopicBuilder) *KafkaBuilder {
	kb.kafkaSetUp.TopicInfo = topicInfo
	return kb
}

func (kb *KafkaBuilder) SetProducerInfo(producerInfo ProducerBuilder) *KafkaBuilder {
	kb.kafkaSetUp.ProducerInfo = producerInfo
	return kb
}

func (kb *KafkaBuilder) SetConsumerGroupInfo(consumerGroupInfo ConsumerGroupBuilder) *KafkaBuilder {
	kb.kafkaSetUp.ConsumerGroup = consumerGroupInfo
	return kb
}

func (kb *KafkaBuilder) Build() Kafka {
	return kb.kafkaSetUp
}
