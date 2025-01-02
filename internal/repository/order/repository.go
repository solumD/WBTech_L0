package order

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/db"
	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/repository"
	"github.com/solumD/WBTech_L0/internal/repository/order/converter"
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
	deliveryID, err := r.createDelivery(ctx, order.Delivery)
	if err != nil {
		return err
	}

	paymentID, err := r.createPayment(ctx, order.Payment)
	if err != nil {
		return err
	}

	itemsIDs, err := r.createItems(ctx, order.Items)
	if err != nil {
		return err
	}

	orderID, err := r.createOrder(ctx, order, deliveryID, paymentID)
	if err != nil {
		return err
	}

	if err := r.createOrdersAndItems(ctx, orderID, itemsIDs); err != nil {
		return err
	}

	return nil
}

// GetOrder gets all orders from storage
func (r *repo) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	orders, err := r.getAllOrders(ctx)
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

	servOrders := []model.Order{}
	for _, o := range orders {
		servOrders = append(servOrders, converter.FromRepoToServiceOrder(o))
	}

	return servOrders, nil
}
