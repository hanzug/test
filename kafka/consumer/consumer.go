// Inventory Service
package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"strings"
)

func main() {
	// Kafka集群的地址
	brokers := []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"}

	// 创建消费者配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建消费者
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic("Failed to start Sarama consumer:" + err.Error())
	}
	defer consumer.Close()

	// 订阅订单主题
	partitionConsumer, err := consumer.ConsumePartition("orders", 0, sarama.OffsetNewest)
	if err != nil {
		panic("Failed to start Sarama partition consumer:" + err.Error())
	}
	defer partitionConsumer.Close()

	// 处理订单
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// 解析订单信息
			orderInfo := strings.Split(string(msg.Value), ",")
			orderID := orderInfo[0]
			productID := orderInfo[1]
			quantity := orderInfo[2]

			// 模拟处理订单并更新库存
			fmt.Printf("Processing order: %s for product: %s, quantity: %s\n", orderID, productID, quantity)

			// 模拟更新库存
			// 此处省略库存更新逻辑

			// 发送库存更新消息
			sendInventoryUpdate(orderID, productID, quantity)
		case err := <-partitionConsumer.Errors():
			fmt.Printf("Failed to consume partition: %s\n", err)
		}
	}
}

func sendInventoryUpdate(orderID, productID, quantity string) {
	// Kafka集群的地址
	brokers := []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"}

	// 创建生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 等待本地复制
	config.Producer.Return.Successes = true            // 成功交付的消息将在success channel返回

	// 连接Kafka集群
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("Failed to start Sarama producer: %s\n", err)
		return
	}
	defer producer.Close()

	// 构造库存更新消息
	message := &sarama.ProducerMessage{
		Topic: "inventory_updates",
		Value: sarama.StringEncoder(fmt.Sprintf("%s,%s,%s", orderID, productID, quantity)),
	}

	// 发送消息
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		fmt.Printf("Failed to send inventory update message: %s\n", err)
	} else {
		fmt.Printf("Inventory update message sent to topic(%s)/partition(%d)/offset(%d)\n", "inventory_updates", partition, offset)
	}
}
