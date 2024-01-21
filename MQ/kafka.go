package MQ

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"sync"
)

func Kafka() {
	// Kafka 生产者示例
	go produceMessages()

	// Kafka 消费者示例
	consumeMessages()
}

func produceMessages() {
	// 设置 Kafka 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	} else {
		zap.L().Info("producer init ok")
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %v", err)
		}
	}()

	// 发送消息到 Kafka 主题
	message := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message to Kafka: %v", err)
	} else {
		zap.L().Info("producer send message ok")
	}

	log.Printf("Message sent to partition %d at offset %d", partition, offset)
}

func consumeMessages() {
	// 设置 Kafka 消费者配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	// 订阅 Kafka 主题
	topics := []string{"test-topic"}
	partitions, err := consumer.Partitions(topics[0])
	if err != nil {
		log.Fatalf("Error retrieving partitions: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(partitions))

	// 处理每个分区的消息
	for _, partition := range partitions {
		go func(partition int32) {
			defer wg.Done()

			partitionConsumer, err := consumer.ConsumePartition(topics[0], partition, sarama.OffsetNewest)
			if err != nil {
				log.Fatalf("Error creating partition consumer: %v", err)
			}
			defer func() {
				if err := partitionConsumer.Close(); err != nil {
					log.Fatalf("Error closing partition consumer: %v", err)
				}
			}()

			for {
				select {
				case msg := <-partitionConsumer.Messages():
					log.Printf("Partition %d - Received message: %s", partition, string(msg.Value))
				case err := <-partitionConsumer.Errors():
					log.Printf("Partition %d - Error: %v", partition, err)
				}
			}
		}(partition)
	}

	// 等待程序终止信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	// 等待所有分区的处理完成
	wg.Wait()
}
