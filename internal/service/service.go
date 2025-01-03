package service

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/model"
)

// OrderService interface of order service
type OrderService interface {
	CreateOrder(ctx context.Context, order model.Order) error
	GetOrderByUID(ctx context.Context, uid string) (model.Order, error)
}
