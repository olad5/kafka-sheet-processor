package consumer

import (
	"context"
	"fmt"

	"github.com/olad5/kafka-sheet-processor/broker"
	"github.com/olad5/kafka-sheet-processor/processor"
)

type Consumer struct {
	messageBroker broker.Broker
	proccessor    *processor.Processor
}

func NewConsumer(messageBroker broker.Broker, processor *processor.Processor) (*Consumer, error) {
	if messageBroker == nil {
		return nil, fmt.Errorf("messageBroker cannot be nil")
	}

	if processor == nil {
		return nil, fmt.Errorf("processor cannot be nil")
	}

	producer := &Consumer{messageBroker, processor}

	return producer, nil
}

func (c *Consumer) Consume(ctx context.Context) error {
	results := make(chan string)
	go c.messageBroker.Consume(ctx, results)
	err := c.proccessor.WriteHeaders()
	if err != nil {
		return fmt.Errorf("Error writing json headers to file: %w", err)
	}

	first := true
	for row := range results {
		err := c.proccessor.Write(row, first)
		if err != nil {
			return fmt.Errorf("Error writing to file: %w", err)
		}
		first = false
	}
	c.proccessor.Flush()

	return nil
}
