package kafkapkg

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
type ConsumerGroupCfg struct {
	Topics  []string
	Brokers []string
	GroupId string
	Config  *sarama.Config
}

type ConsumerGroupCfgBuilder struct {
	consumerCfg ConsumerGroupCfg
}

func NewConsumerGroupCfgBuilder() *ConsumerGroupCfgBuilder {
	return &ConsumerGroupCfgBuilder{consumerCfg: ConsumerGroupCfg{Config: sarama.NewConfig(), Brokers: make([]string, 0), Topics: make([]string, 0)}}
}

func (cb *ConsumerGroupCfgBuilder) SetConsumerGroupRebalanceStratagy(strategy sarama.BalanceStrategy) *ConsumerGroupCfgBuilder {
	cb.consumerCfg.Config.Consumer.Group.Rebalance.Strategy = strategy
	return cb
}

func (cb *ConsumerGroupCfgBuilder) SetConsumerOffset(offsetVal int64) *ConsumerGroupCfgBuilder {
	cb.consumerCfg.Config.Consumer.Offsets.Initial = offsetVal
	return cb
}

func (cb *ConsumerGroupCfgBuilder) SetTopics(topics []string) *ConsumerGroupCfgBuilder {
	cb.consumerCfg.Topics = append(cb.consumerCfg.Topics, topics...)
	return cb
}

func (cb *ConsumerGroupCfgBuilder) SetBrokers(brokers []string) *ConsumerGroupCfgBuilder {
	cb.consumerCfg.Brokers = append(cb.consumerCfg.Brokers, brokers...)
	return cb
}

func (cb *ConsumerGroupCfgBuilder) SetGroupId(groupId string) *ConsumerGroupCfgBuilder {
	cb.consumerCfg.GroupId = groupId
	return cb
}

func (cb *ConsumerGroupCfgBuilder) Build() ConsumerGroupCfg {
	return cb.consumerCfg
}
