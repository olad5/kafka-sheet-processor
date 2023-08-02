package broker

import (
	"bufio"
	"context"
)

type Broker interface {
	Consume(ctx context.Context, results chan string)
	Produce(ctx context.Context, scanner *bufio.Scanner)
}
