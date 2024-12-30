package order

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/client/db"
	modelRepo "github.com/solumD/WBTech_L0/internal/repository/order/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) getOrders(ctx context.Context) ([]modelRepo.Order, error) {
	query, args, err := sq.Select("*").From(ordersTableName).ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "order_repository.getOrders",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []modelRepo.Order{}
	for rows.Next() {
		var order modelRepo.Order

		err := rows.Scan(
			&order.Id,
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.DeliveryId,
			&order.PaymentId,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
