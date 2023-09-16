package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
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
	config.Producer.Return.Successes = true                  // if true, must read Successes channel
	config.Producer.Return.Errors = true                     // default enabled, must read Errors channel

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

	log.Println("Producer starts. Press CTRL+C to exit.")

outer:
	for i := 0; ; {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
		}

		select {
		case producer.Input() <- message:
			i++
			log.Printf("Produced message: %s", message.Value)
		case successMsg := <-producer.Successes():
			log.Printf("Producer success: %v", successMsg)
			fmt.Println("== topic:", successMsg.Topic)
			fmt.Println("== value:", successMsg.Value)
			fmt.Println("== partition:", successMsg.Partition)
			fmt.Println("== offset:", successMsg.Offset)
		case err := <-producer.Errors():
			log.Printf("Producer error: %v", err)
		case <-signals:
			log.Println("Signal interrupt.")
			break outer
		}
		time.Sleep(1 * time.Second)
	}

	return
}
