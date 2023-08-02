package main

import (
	"context"
	"log"

	"github.com/olad5/kafka-sheet-processor/broker/kafka"
	"github.com/olad5/kafka-sheet-processor/config"
	"github.com/olad5/kafka-sheet-processor/consumer"
	"github.com/olad5/kafka-sheet-processor/processor"
)

func main() {
	ctx := context.Background()

	messageBroker, err := kafka.NewKafkaBroker(config.KAFKA_HOST)
	if err != nil {
		log.Fatal("failed to Initialize Kafka broker: ", err)
	}

	processor, err := processor.NewProcessor("result.json")
	if err != nil {
		log.Fatal("failed to Initialize Processor: ", err)
	}

	consumer, err := consumer.NewConsumer(messageBroker, processor)
	if err != nil {
		log.Fatal("failed to Initialize Consumer : ", err)
	}

	err = consumer.Consume(ctx)
	if err != nil {
		log.Fatal("Producer error", err)
	}
}
