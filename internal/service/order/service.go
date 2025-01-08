package order

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/solumD/WBTech_L0/internal/cache"
	"github.com/solumD/WBTech_L0/internal/db"
	"github.com/solumD/WBTech_L0/internal/logger"
	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/repository"
	"github.com/solumD/WBTech_L0/internal/service"

	"github.com/IBM/sarama"
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
		logger.Error("failed to save order in cache", zap.Error(err))
		return errCache
	}

	logger.Info("saved order in cache", zap.Any("order", order))

	return nil
}

// GetOrderByUID gets order from cache by id
func (s *srv) GetOrderByUID(_ context.Context, uid string) (model.Order, error) {
	order, errCache := s.orderCache.GetOrderByUID(uid)
	if errCache != nil {
		logger.Error("failed to ger order from cache", zap.Error(errCache))

		return model.Order{}, errCache
	}

	logger.Info("got order from cache", zap.Any("order", order))
	return order, nil
}

// ConsumeOrders gets kafka order-messages from orders chan and saves them in repository
func (s *srv) ConsumeOrders(ctx context.Context, orders chan *sarama.ConsumerMessage) error {
	if orders == nil {
		return fmt.Errorf("orders chan is nil")
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-orders:
				if !ok {
					logger.Info("orders chan is closed, stopping service consumer")
					return
				}

				order := &model.Order{}
				if err := json.Unmarshal(msg.Value, order); err != nil {
					logger.Error("failed to unmarshal order", zap.Error(err))

				}

				logger.Info("recieved order from consumer", zap.Any("order", order))

				if err := s.CreateOrder(ctx, *order); err != nil {
					logger.Error("failed to create order", zap.Error(err))
				}
			}

		}
	}()

	wg.Wait()

	return nil
}
