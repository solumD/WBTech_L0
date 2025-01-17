package inmemory

import (
	"context"
	"errors"
	"sync"

	"github.com/solumD/WBTech_L0/internal/cache"
	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/repository"
)

var (
	errAlreadyExists = errors.New("order already exists in cache")
	errDoesntExist   = errors.New("order doesn't exist in cache")
)

type inMemoryOrderCache struct {
	orders map[string]model.Order
	mu     *sync.RWMutex
}

// New returns new in-memory order cache object
func New() cache.OrderCache {
	return &inMemoryOrderCache{
		orders: make(map[string]model.Order),
		mu:     &sync.RWMutex{},
	}
}

// SaveOrder saves order in cache by its' uid
func (c *inMemoryOrderCache) SaveOrder(uid string, order model.Order) error {
	c.mu.RLock()
	_, exist := c.orders[uid]
	c.mu.RUnlock()

	if exist {
		return errAlreadyExists
	}

	c.mu.Lock()
	c.orders[uid] = order
	c.mu.Unlock()

	return nil
}

// GetOrderByUID gets order from cache by its' uid
func (c *inMemoryOrderCache) GetOrderByUID(uid string) (model.Order, error) {
	c.mu.RLock()
	order, exist := c.orders[uid]
	c.mu.RUnlock()

	if !exist {
		return model.Order{}, errDoesntExist
	}

	return order, nil
}

// LoadOrders loads all orders from repository
func (c *inMemoryOrderCache) LoadOrders(ctx context.Context, repo repository.OrderRepository) error {
	orders, err := repo.GetAllOrders(ctx)
	if err != nil {
		return err
	}

	c.mu.Lock()
	for _, o := range orders {
		c.orders[o.OrderUID] = o
	}
	c.mu.Unlock()

	return nil
}
