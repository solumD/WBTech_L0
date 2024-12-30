package order

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/client/db"
	modelRepo "github.com/solumD/WBTech_L0/internal/repository/order/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) getAndSetAllOrders(ctx context.Context) ([]modelRepo.Order, error) {
	query, args, err := sq.Select("*").From(ordersTableName).ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "order_repository.getAndSetOrders",
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

func (r *repo) getAndSetAllDelivery(ctx context.Context, orders []modelRepo.Order) ([]modelRepo.Order, error) {
	for i, order := range orders {
		query, args, err := sq.Select("*").
			From(deliveryTableName).
			PlaceholderFormat(sq.Dollar).
			Where(sq.Eq{idCol: order.DeliveryId}).
			ToSql()

		if err != nil {
			return nil, err
		}

		q := db.Query{
			Name:     "order_repository.getAndSetAllDelivery",
			QueryRaw: query,
		}

		var delivery modelRepo.Delivery
		if err = r.db.DB().ScanOneContext(ctx, &delivery, q, args...); err != nil {
			return nil, err
		}

		orders[i].Delivery = delivery
	}

	return orders, nil
}

func (r *repo) getAndSetAllPayment(ctx context.Context, orders []modelRepo.Order) ([]modelRepo.Order, error) {
	for i, order := range orders {
		query, args, err := sq.Select("*").
			From(paymentTableName).
			PlaceholderFormat(sq.Dollar).
			Where(sq.Eq{idCol: order.PaymentId}).
			ToSql()

		if err != nil {
			return nil, err
		}

		q := db.Query{
			Name:     "order_repository.getAndSetAllPayment",
			QueryRaw: query,
		}

		var payment modelRepo.Payment
		if err = r.db.DB().ScanOneContext(ctx, &payment, q, args...); err != nil {
			return nil, err
		}

		orders[i].Payment = payment
	}

	return orders, nil
}

func (r *repo) getAndSetAllItems(ctx context.Context, orders []modelRepo.Order) ([]modelRepo.Order, error) {
	for i, order := range orders {
		query, args, err := sq.Select(itemIdCol).
			From(ordersAndItemsTableName).
			PlaceholderFormat(sq.Dollar).
			Where(sq.Eq{orderIdCol: order.Id}).
			ToSql()

		if err != nil {
			return nil, err
		}

		q := db.Query{
			Name:     "order_repository.getAndSetAllItems",
			QueryRaw: query,
		}

		rows, err := r.db.DB().QueryContext(ctx, q, args...)
		if err != nil {

			return nil, err
		}
		defer rows.Close()

		ids := []int{}
		for rows.Next() {
			var id int

			err := rows.Scan(&id)
			if err != nil {
				return nil, err
			}

			ids = append(ids, id)
		}

		items := []modelRepo.Item{}
		for _, id := range ids {
			query, args, err = sq.Select("*").
				From(itemTableName).
				PlaceholderFormat(sq.Dollar).
				Where(sq.Eq{idCol: id}).
				ToSql()

			if err != nil {
				return nil, err
			}

			q = db.Query{
				Name:     "order_repository.getAndSetAllItems",
				QueryRaw: query,
			}

			var item modelRepo.Item
			if err = r.db.DB().ScanOneContext(ctx, &item, q, args...); err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		orders[i].Items = items
	}

	return orders, nil
}
