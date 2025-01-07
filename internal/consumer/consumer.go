package consumer

import (
	"context"

	"github.com/IBM/sarama"
)

// OrderConsumer interface of order consumer
type OrderConsumer interface {
	Consume(ctx context.Context) (chan *sarama.ConsumerMessage, error)
}
