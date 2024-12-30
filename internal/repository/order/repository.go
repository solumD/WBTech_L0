package order

import (
	"context"
	"log"

	"github.com/solumD/WBTech_L0/internal/client/db"
	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/repository"
)

type repo struct {
	db db.Client
}

// New returns new repository object
func New(db db.Client) repository.OrderRepository {
	return &repo{
		db: db,
	}
}

// CreateOrder inserts order in storage
func (r *repo) CreateOrder(ctx context.Context, order model.Order) error {
	deliveryId, err := r.createDelivery(ctx, order.Delivery)
	if err != nil {
		return err
	}

	paymentId, err := r.createPayment(ctx, order.Payment)
	if err != nil {
		return err
	}

	itemsIds, err := r.createItems(ctx, order.Items)
	if err != nil {
		return err
	}

	orderId, err := r.createOrder(ctx, order, deliveryId, paymentId)
	if err != nil {
		return err
	}

	if err := r.createOrdersAndItems(ctx, orderId, itemsIds); err != nil {
		return err
	}

	return nil
}

// GetOrder gets all orders from storage
func (r *repo) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.getAndSetAllOrders(ctx)
	if err != nil {
		return nil, err
	}

	orders, err = r.getAndSetAllDelivery(ctx, orders)
	if err != nil {
		return nil, err
	}

	orders, err = r.getAndSetAllPayment(ctx, orders)
	if err != nil {
		return nil, err
	}

	orders, err = r.getAndSetAllItems(ctx, orders)
	if err != nil {
		return nil, err
	}

	for _, o := range orders {
		log.Println(o)
	}

	return nil, nil
}
