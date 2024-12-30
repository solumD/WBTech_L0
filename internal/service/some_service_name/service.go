package someservicename

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/client/db"
	"github.com/solumD/WBTech_L0/internal/repository"
	"github.com/solumD/WBTech_L0/internal/service"
)

type srv struct {
	orderRepository repository.OrderRepository
	txManager       db.TxManager
}

// New returns new service object
func New(orderRepository repository.OrderRepository, txManager db.TxManager) service.SomeService {
	return &srv{
		orderRepository: orderRepository,
		txManager:       txManager,
	}
}

// SomeMethod ...
func (s *srv) SomeMethod(_ context.Context, _ ...interface{}) (interface{}, error) {
	// some business logic
	return struct{}{}, nil
}
