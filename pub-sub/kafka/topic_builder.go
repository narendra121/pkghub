package kafka

import "github.com/IBM/sarama"

/*
------------------------------------Usage-------------------------

	tb := kafka.NewTopicBuilder()
	tb.
		SetBrokers([]string{"kafka:9092"}).
		SetPartitions(3).
		SetReplicationFactor(1).
		SetTopicName("test3").
		Build()

	kb := kafka.NewKafkaBuilder()
	kb.SetTopicInfo(*tb).Build

    kf := kafka.NewPubSubFactory(kb)
    err := kf.CreateNewTopic()
*/

type Topic struct {
	TopicName   string
	Brokers     []string
	TopicConfig *sarama.TopicDetail
}

type TopicBuilder struct {
	topic Topic
}

func NewTopicBuilder() *TopicBuilder {
	return &TopicBuilder{topic: Topic{Brokers: make([]string, 0), TopicConfig: &sarama.TopicDetail{ConfigEntries: map[string]*string{}}}}
}

func (t *TopicBuilder) SetPartitions(partitions int32) *TopicBuilder {
	t.topic.TopicConfig.NumPartitions = partitions
	return t
}

func (t *TopicBuilder) SetReplicationFactor(repllicationFactor int16) *TopicBuilder {
	t.topic.TopicConfig.ReplicationFactor = repllicationFactor
	return t
}

func (t *TopicBuilder) SetTopicName(topicName string) *TopicBuilder {
	t.topic.TopicName = topicName
	return t
}

func (t *TopicBuilder) SetConfigEntries(key, val string) *TopicBuilder {
	t.topic.TopicConfig.ConfigEntries[key] = &val
	return t
}

func (t *TopicBuilder) SetBrokers(brokers []string) *TopicBuilder {
	t.topic.Brokers = append(t.topic.Brokers, brokers...)
	return t
}

func (t *TopicBuilder) Build() Topic {
	return t.topic
}
