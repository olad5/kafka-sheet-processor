package main

import (
	"context"
	"log"

	"github.com/olad5/kafka-sheet-processor/broker/kafka"
	"github.com/olad5/kafka-sheet-processor/config"
	"github.com/olad5/kafka-sheet-processor/producer"
)

func main() {
	ctx := context.Background()

	messageBroker, err := kafka.NewKafkaBroker(config.KAFKA_HOST)
	if err != nil {
		log.Fatal("failed to Initialize Kafka broker: ", err)
	}

	producer, err := producer.NewProducer("results.csv", messageBroker)
	if err != nil {
		log.Fatal("failed to Initialize Producer : ", err)
	}

	err = producer.Produce(ctx)
	if err != nil {
		log.Fatal("Producer error", err)
	}
}
