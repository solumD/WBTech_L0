package cache

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/repository"
)

// OrderCache interface of order cache
type OrderCache interface {
	SaveOrder(uid string, order *model.Order) error
	GetOrderByUID(uid string) (*model.Order, error)
	LoadOrders(ctx context.Context, repo repository.OrderRepository) error
}
