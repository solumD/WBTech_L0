package someservicename

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/cache"
	"github.com/solumD/WBTech_L0/internal/db"
	"github.com/solumD/WBTech_L0/internal/repository"
	"github.com/solumD/WBTech_L0/internal/service"
)

type srv struct {
	orderRepository repository.OrderRepository
	orderCache      cache.OrderCache
	txManager       db.TxManager
}

// New returns new service object
func New(orderRepository repository.OrderRepository, orderCache cache.OrderCache, txManager db.TxManager) service.SomeService {
	return &srv{
		orderRepository: orderRepository,
		orderCache:      orderCache,
		txManager:       txManager,
	}
}

// SomeMethod ...
func (s *srv) SomeMethod(_ context.Context, _ ...interface{}) (interface{}, error) {
	// some business logic
	return struct{}{}, nil
}
