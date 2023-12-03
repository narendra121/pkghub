package kafka

import "github.com/IBM/sarama"

/*
----------------------------Usage-------------------------------

cgb:= kafka.NewConsumerGroupBuilder()

	cgb.SetBrokers([]string{"kafka:9092"}).
		SetConsumerGroupRebalanceStratagy(sarama.NewBalanceStrategyRoundRobin()).
		SetConsumerOffset(sarama.OffsetOldest).
		SetTopics([]string{"test3"}).
		SetGroupId("test3").
		Build()

		kb := kafka.NewKafkaBuilder()
		kb.SetConsumerGroupInfo(*cgb).Build()

		kf := kafka.NewPubSubFactory(kb)

	cgroup, err := kf.CreateConsumerGroup()

handler := &kafka.MessageHandler{CustomMessageHandler: customMessageHandler}

	func customMessageHandler(msg *sarama.ConsumerMessage) {
		log.Infof("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(msg.Value), msg.Timestamp, msg.Topic)

}
kf.AddConsumerToConsumerGroup(cgroup, handler)

	kf.AddConsumerToConsumerGroup(cgroup)
*/
type ConsumerGroup struct {
	Topics  []string
	Brokers []string
	GroupId string
	Config  *sarama.Config
}

type ConsumerGroupBuilder struct {
	consumerCfg ConsumerGroup
}

func NewConsumerGroupBuilder() *ConsumerGroupBuilder {
	return &ConsumerGroupBuilder{consumerCfg: ConsumerGroup{Config: sarama.NewConfig(), Brokers: make([]string, 0), Topics: make([]string, 0)}}
}

func (cb *ConsumerGroupBuilder) SetConsumerGroupRebalanceStratagy(strategy sarama.BalanceStrategy) *ConsumerGroupBuilder {
	cb.consumerCfg.Config.Consumer.Group.Rebalance.Strategy = strategy
	return cb
}

func (cb *ConsumerGroupBuilder) SetConsumerOffset(offsetVal int64) *ConsumerGroupBuilder {
	cb.consumerCfg.Config.Consumer.Offsets.Initial = offsetVal
	return cb
}

func (cb *ConsumerGroupBuilder) SetTopics(topics []string) *ConsumerGroupBuilder {
	cb.consumerCfg.Topics = append(cb.consumerCfg.Topics, topics...)
	return cb
}

func (cb *ConsumerGroupBuilder) SetBrokers(brokers []string) *ConsumerGroupBuilder {
	cb.consumerCfg.Brokers = append(cb.consumerCfg.Brokers, brokers...)
	return cb
}

func (cb *ConsumerGroupBuilder) SetGroupId(groupId string) *ConsumerGroupBuilder {
	cb.consumerCfg.GroupId = groupId
	return cb
}

func (cb *ConsumerGroupBuilder) Build() ConsumerGroup {
	return cb.consumerCfg
}
