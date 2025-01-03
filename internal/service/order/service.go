package order

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/cache"
	"github.com/solumD/WBTech_L0/internal/db"
	"github.com/solumD/WBTech_L0/internal/logger"
	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/repository"
	"github.com/solumD/WBTech_L0/internal/service"
	"go.uber.org/zap"
)

type srv struct {
	orderRepository repository.OrderRepository
	orderCache      cache.OrderCache
	txManager       db.TxManager
}

// New returns new order service object
func New(orderRepository repository.OrderRepository, orderCache cache.OrderCache, txManager db.TxManager) service.OrderService {
	return &srv{
		orderRepository: orderRepository,
		orderCache:      orderCache,
		txManager:       txManager,
	}
}

// CreateOrder creates order in repository and saves it in cache
func (s *srv) CreateOrder(ctx context.Context, order model.Order) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.orderRepository.CreateOrder(ctx, order)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	logger.Info("saved order in repository", zap.Any("order", order))

	errCache := s.orderCache.SaveOrder(order.OrderUID, order)
	if errCache != nil {
		return errCache
	}

	logger.Info("saved order in cache", zap.Any("order", order))

	return nil
}

// GetOrderByUID gets order from cache by id
func (s *srv) GetOrderByUID(_ context.Context, uid string) (model.Order, error) {
	order, errCache := s.orderCache.GetOrderByUID(uid)
	if errCache != nil {
		return model.Order{}, errCache
	}

	logger.Info("got order from cache", zap.Any("order", order))
	return order, nil
}
