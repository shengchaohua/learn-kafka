package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"

	"learn-kafka/examples/single_node/conf"
)

func main() {
	broker := conf.KafkaHost
	topic := conf.FirstTopic

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // 等待 Leader 发送成功
	config.Producer.Compression = sarama.CompressionSnappy   // 使用 Snappy 压缩
	config.Producer.Flush.Frequency = 500 * time.Millisecond // 定期刷新

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("Error closing producer: %v", err)
		}
	}()

	for i := 0; i < 10; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
		}

		producer.Input() <- message
		log.Printf("Produced message: %s", message.Value)
		time.Sleep(1 * time.Second)
	}

	log.Println("END")
}
