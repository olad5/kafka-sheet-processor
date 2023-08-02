package producer

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/olad5/kafka-sheet-processor/broker"
)

type Producer struct {
	messageBroker broker.Broker
	file          *os.File
}

func NewProducer(filepath string, messageBroker broker.Broker) (*Producer, error) {
	if filepath == "" {
		return nil, fmt.Errorf("filepath cannot be empty")
	}
	if messageBroker == nil {
		return nil, fmt.Errorf("messageBroker cannot be nil")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file :%w", err)
	}

	producer := &Producer{messageBroker, file}

	return producer, nil
}

func (p *Producer) Produce(ctx context.Context) error {
	scanner := bufio.NewScanner(p.file)
	scanner.Scan()
	p.messageBroker.Produce(ctx, scanner)
	err := p.file.Close()
	if err != nil {
		return fmt.Errorf("Error closing file: %w", err)
	}
	return nil
}
