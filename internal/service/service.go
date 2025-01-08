package service

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/model"

	"github.com/IBM/sarama"
)

// OrderService interface of order service
type OrderService interface {
	CreateOrder(ctx context.Context, order model.Order) error
	GetOrderByUID(ctx context.Context, uid string) (model.Order, error)
	ConsumeOrders(ctx context.Context, orders chan *sarama.ConsumerMessage) error
}
