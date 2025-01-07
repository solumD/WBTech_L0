package service

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/solumD/WBTech_L0/internal/model"
)

// OrderService interface of order service
type OrderService interface {
	Consume(ctx context.Context, msgIn chan *sarama.ConsumerMessage)
	CreateOrder(ctx context.Context, order model.Order) error
	GetOrderByUID(ctx context.Context, uid string) (model.Order, error)
}
