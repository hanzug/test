package MQ

import "testing"

func TestKafka(t *testing.T) {
	// Kafka 生产者示例
	go produceMessages()

	// Kafka 消费者示例
	consumeMessages()
}
