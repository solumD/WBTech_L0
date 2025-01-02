package repository

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/model"
)

// OrderRepository interface of orders repository
type OrderRepository interface {
	CreateOrder(ctx context.Context, order model.Order) error
	GetAllOrders(ctx context.Context) ([]model.Order, error)
}
