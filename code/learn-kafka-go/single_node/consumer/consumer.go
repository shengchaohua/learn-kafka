package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"

	"learn-kafka-go/single_node/conf"
)

func Consumer() {
	broker := conf.KafkaHost
	topic := conf.FirstTopic

	config := sarama.NewConfig()
	// conf.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的偏移量开始消费，增量消费
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

	consumerPartition, err := consumer.ConsumePartition(topic, 0, 0)
	if err != nil {
		log.Fatalf("Error from consumer: %v", err)
	}

	go func() {
		for {
			select {
			case message := <-consumerPartition.Messages():
				log.Printf("Consume message:")
				log.Print("== topic:", message.Topic)
				log.Print("== value:", string(message.Value))
				log.Print("== partition:", message.Partition)
				log.Print("== offset:", message.Offset)
			case err := <-consumerPartition.Errors():
				log.Println(err)
			}
		}
	}()

	<-signals
	log.Println("Signal interrupt.")
}
