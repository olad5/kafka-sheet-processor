package kafka

import (
	"bufio"
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/olad5/kafka-sheet-processor/config"
)

type KafkaBroker struct {
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func NewKafkaBroker(kafkaHost string) (*KafkaBroker, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{kafkaHost}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewSyncProducer, err :%w", err)
	}
	consumer, err := sarama.NewConsumer([]string{kafkaHost}, config)
	if err != nil {
		producer.Close()
		return nil, fmt.Errorf("failed to initialize NewConsumer, err :%w", err)
	}
	return &KafkaBroker{producer, consumer}, nil
}

func (k *KafkaBroker) Produce(ctx context.Context, scanner *bufio.Scanner) {
	for scanner.Scan() {
		msg := &sarama.ProducerMessage{Topic: config.KAFKA_TOPIC, Key: nil, Value: sarama.StringEncoder(scanner.Text())}
		_, _, err := k.producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
	}
	sentinelValue := ""
	k.producer.SendMessage(&sarama.ProducerMessage{Topic: config.KAFKA_TOPIC, Key: nil, Value: sarama.StringEncoder(sentinelValue)})
	err := k.producer.Close()
	if err != nil {
		fmt.Errorf("Error closing Kafka Producer :%w", err)
		return
	}
}

func (k *KafkaBroker) Consume(ctx context.Context, results chan string) {
	partitionConsumer, err := k.consumer.ConsumePartition(config.KAFKA_TOPIC, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}

	for {
		select {
		case message, ok := <-partitionConsumer.Messages():
			if ok {
				messageValue := string(message.Value)
				if messageValue == "" {
					err = partitionConsumer.Close()
					if err != nil {
						log.Println("partitionConsumer close err: ", err)
						return
					}

					err = k.consumer.Close()
					if err != nil {
						log.Println("Consumer close err: ", err)
						return
					}
					break
				}
				results <- messageValue
			} else {
				close(results)
				return
			}
		}
	}
}
