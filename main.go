package main

import (
	"fmt"

	"github.com/narendra121/pkghub/auth"
)

func main() {

	// jb := auth.NewJwtBuilder().AddEmail("hello").AddSignInSalt("fffffff").AddTokenExpiry(2).AddUserName("naren").Build()
	// fac := auth.NewTokenFactory(jb).GenerateSignedToken()
	// fact := auth.NewTokenFactory(jb)
	// t := fact.GenerateSignedToken()
	// v := fact.IsTokenValid("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvIiwic2lnbmluX3NhbHQiOiJmZmZmZmZmIiwidG9rZW5fZXhwIjoxNzAxMzg0MjE3LCJ1c2VyX25hbWUiOiJuYXJlbiJ9.RDY5qfaF-cfLDk5caaNWkwLUgGPKN1GkZZYI2xNYqqw")
	// r := fact.RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvIiwic2lnbmluX3NhbHQiOiJmZmZmZmZmIiwidG9rZW5fZXhwIjoxNzAxMzg0MTg4LCJ1c2VyX25hbWUiOiJuYXJlbiJ9.TbReY6zOkm01rkdm16H9uP5-sVieKRB-pKZORxqppdg")
	// fmt.Println("token", fac)
	// fmt.Println(t)
	// fmt.Println(v)
	// fmt.Println(r)

	// tb := kafka.NewTopicBuilder()
	// tb.
	// 	SetBrokers([]string{"kafka:9092"}).
	// 	SetPartitions(3).
	// 	SetReplicationFactor(1).
	// 	SetTopicName("naren").
	// 	Build()

	// pb := kafka.NewProducerBuilder()
	// pb.SetBrokers([]string{"kafka:9092"}).
	// 	SetAckType(sarama.WaitForAll).
	// 	Build()

	// cgb := kafka.NewConsumerGroupBuilder()
	// cgb.SetBrokers([]string{"kafka:9092"}).
	// 	SetConsumerGroupRebalanceStratagy(sarama.NewBalanceStrategyRoundRobin()).
	// 	SetConsumerOffset(sarama.OffsetOldest).
	// 	SetTopics([]string{"naren", "test3"}).
	// 	SetGroupId("test6").
	// 	Build()

	// kb := kafka.NewKafkaBuilder()

	// kb.SetProducerInfo(*pb).
	// 	SetTopicInfo(*tb).
	// 	SetConsumerGroupInfo(*cgb).
	// 	Build()

	// kf := pubsub.NewPubSubFactory(kb)
	// err := kf.CreateNewTopic()
	// log.Println(err)

	// producer, err := kf.CreateSyncProducer()
	// log.Println(err)

	// kf.SendMessage(producer, "test3", "my test")
	// kf.SendMessage(producer, "naren", "mysssd")

	// handler := &kafka.MessageHandler{CustomMessageHandler: customMessageHandler}
	// cgroup, err := kf.CreateConsumerGroup()
	// log.Println("cg  -----                ", err)
	// kf.AddConsumerToConsumerGroup(cgroup, handler)

	// ppstb := postgresdb.NewPostgresDbBuilder().
	// 	SetUser("narendra").
	// 	SetPassword("123456").
	// 	SetHost("localhost").
	// 	SetDbPort("5432").
	// 	SetDbName("practice").Build()
	// dbf := db.NewDbFactory(&ppstb)
	// dbm, _ := dbf.Connect()
	// r := dbm.(*gorm.DB)
	// r.Migrator().CreateTable(&UserJam{})
	// _ = dbf.Connect()
	// dbf.CreateTable(&UserJam{})
	jb := auth.NewJwtBuilder().AddUserName("1234567").Build()
	tk := auth.NewTokenFactory(&jb)
	// t := tk.GenerateSignedToken(2, "sss", nil)
	// fmt.Println(t)
	v := tk.IsTokenValid("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl9leHAiOjE3MDIxMTY5ODAsInVzZXJfbmFtZSI6IjEyMzQ1NjcifQ.uitDNgeoq26jIeTwCBKlfVCw2bgNCyPRdexWfNyT7-c", "sss", nil)
	fmt.Println(v)
	// r := tk.RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl9leHAiOjE3MDIxMTY5MDQsInVzZXJfbmFtZSI6IjEyMzQ1NjcifQ.jzwUOkpIDOJQDhCUq-FQgLTQ4cYG1h98Uru_xWP1OHs", "sss", 2, nil)
	// fmt.Println(r)
}

// type UserJam struct {
// 	gorm.Model
// 	UserName string
// 	Email    string
// }

// func customMessageHandler(msg *sarama.ConsumerMessage) {
// 	log.Infof("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(msg.Value), msg.Timestamp, msg.Topic)

// }
