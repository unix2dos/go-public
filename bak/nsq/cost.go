package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "levonfly", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 处理信息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))

	// 连接 nsqd 的 tcp 连接
	if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
