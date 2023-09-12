package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"

	"learn-kafka/examples/first_topic/conf"
)

func main() {
	broker := conf.KafkaHost
	topic := conf.FirstTopic

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // 等待所有副本都保存成功
	config.Producer.Compression = sarama.CompressionSnappy   // 使用Snappy压缩
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

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 使用goroutine发送消息
	go func() {
		for i := 0; ; i++ {
			message := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
			}

			select {
			case producer.Input() <- message:
				log.Printf("Produced message: %s", message.Value)
			case <-signals:
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	<-signals
	log.Println("Signal interrupt.")
}
