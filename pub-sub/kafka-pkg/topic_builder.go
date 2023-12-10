package kafkapkg

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

type TopicCfg struct {
	TopicName   string
	Brokers     []string
	TopicConfig *sarama.TopicDetail
}

type TopicCfgBuilder struct {
	topicCfg TopicCfg
}

func NewTopicBuilder() *TopicCfgBuilder {
	return &TopicCfgBuilder{topicCfg: TopicCfg{Brokers: make([]string, 0), TopicConfig: &sarama.TopicDetail{ConfigEntries: map[string]*string{}}}}
}

func (t *TopicCfgBuilder) SetPartitions(partitions int32) *TopicCfgBuilder {
	t.topicCfg.TopicConfig.NumPartitions = partitions
	return t
}

func (t *TopicCfgBuilder) SetReplicationFactor(repllicationFactor int16) *TopicCfgBuilder {
	t.topicCfg.TopicConfig.ReplicationFactor = repllicationFactor
	return t
}

func (t *TopicCfgBuilder) SetTopicName(topicName string) *TopicCfgBuilder {
	t.topicCfg.TopicName = topicName
	return t
}

func (t *TopicCfgBuilder) SetConfigEntries(key, val string) *TopicCfgBuilder {
	t.topicCfg.TopicConfig.ConfigEntries[key] = &val
	return t
}

func (t *TopicCfgBuilder) SetBrokers(brokers []string) *TopicCfgBuilder {
	t.topicCfg.Brokers = append(t.topicCfg.Brokers, brokers...)
	return t
}

func (t *TopicCfgBuilder) Build() TopicCfg {
	return t.topicCfg
}
