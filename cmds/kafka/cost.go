package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	// 取出老的值
	oldest, _ := consumer.ConsumePartition("web_log", 0, sarama.OffsetOldest)
	defer oldest.AsyncClose()
	for msg := range oldest.Messages() {
		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}

	// 消费新增加的值
	newest, _ := consumer.ConsumePartition("web_log", 0, sarama.OffsetNewest)
	defer newest.AsyncClose()
	for msg := range newest.Messages() {
		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
