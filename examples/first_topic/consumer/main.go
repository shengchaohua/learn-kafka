package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"

	"learn-kafka/examples/first_topic/conf"
)

func main() {
	broker := conf.KafkaHost
	topic := conf.FirstTopic

	config := sarama.NewConfig()
	// config.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的偏移量开始消费，增量消费
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 从最早的偏移量开始消费

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("Error closing consumer: %v", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumerPartition, err := consumer.ConsumePartition(topic, 0, 10)
	if err != nil {
		log.Fatalf("Error from consumer: %v", err)
	}

	go func() {
		for {
			select {
			case message := <-consumerPartition.Messages():
				log.Printf("Consume message:")
				fmt.Println("== topic:", message.Topic)
				fmt.Println("== value:", string(message.Value))
				fmt.Println("== partition:", message.Partition)
				fmt.Println("== offset:", message.Offset)
			case err := <-consumerPartition.Errors():
				log.Println(err)
			}
		}
	}()

	<-signals
	log.Println("Signal interrupt.")
}

type ConsumerHandler struct{}

func (h *ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Consumed message: %s\n", message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}
