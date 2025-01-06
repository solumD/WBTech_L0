package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/solumD/WBTech_L0/internal/closer"
	"github.com/solumD/WBTech_L0/internal/consumer"
)

const (
	ordersTopicName = "orders-topic"
)

type orderConsumer struct {
	consumer  sarama.Consumer
	topicName string
}

// NewOrderConsumer returns new order consumer
func NewOrderConsumer(brokers []string) (consumer.OrderConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &orderConsumer{
		consumer:  consumer,
		topicName: ordersTopicName,
	}, nil
}

// Consume recieves chan for producer messages,
// consumes messages and sends them in msgOut
func (c *orderConsumer) Consume(ctx context.Context, msgOut chan *sarama.ConsumerMessage) error {
	pc, err := c.consumer.ConsumePartition(c.topicName, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	closer.Add(pc.Close)

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-pc.Messages():
			if !ok {
				return nil
			}

			msgOut <- msg
		}
	}

}
