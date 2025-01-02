package order

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/db"
	"github.com/solumD/WBTech_L0/internal/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) createDelivery(ctx context.Context, delivery model.Delivery) (int, error) {
	query, args, err :=
		sq.Insert(deliveryTableName).
			PlaceholderFormat(sq.Dollar).
			Columns(
				delNameCol,
				phoneCol,
				zipCol,
				cityCol,
				addressCol,
				regionCol,
				emailCol,
			).
			Values(
				delivery.Name,
				delivery.Phone,
				delivery.Zip,
				delivery.City,
				delivery.Address,
				delivery.Region,
				delivery.Email,
			).
			Suffix("RETURNING id").ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "order_repository.createDelivery",
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) createPayment(ctx context.Context, payment model.Payment) (int, error) {
	query, args, err :=
		sq.Insert(paymentTableName).
			PlaceholderFormat(sq.Dollar).
			Columns(
				transactionCol,
				requestIDCol,
				currencyCol,
				providerCol,
				amountCol,
				paymentDtCol,
				bankCol,
				deliveryCostCol,
				goodsTotalCol,
				customFeeCol,
			).
			Values(
				payment.Transaction,
				payment.RequestID,
				payment.Currency,
				payment.Provider,
				payment.Amount,
				payment.PaymentDt,
				payment.Bank,
				payment.DeliveryCost,
				payment.GoodsTotal,
				payment.CustomFee,
			).
			Suffix("RETURNING id").ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "order_repository.createPayment",
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) createItems(ctx context.Context, items []model.Item) ([]int, error) {
	itemsIDs := make([]int, 0, 10)
	for _, item := range items {
		query, args, err :=
			sq.Insert(itemTableName).
				PlaceholderFormat(sq.Dollar).
				Columns(
					chrtIDCol,
					itemsTrackNumberCol,
					priceCol,
					ridCol,
					itemNameCol,
					saleCol,
					sizeCol,
					totalPriceCol,
					nmIDCol,
					brandCol,
					statusCol,
				).
				Values(
					item.ChrtID,
					item.TrackNumber,
					item.Price,
					item.Rid,
					item.Name,
					item.Sale,
					item.Size,
					item.TotalPrice,
					item.NmID,
					item.Brand,
					item.Status,
				).
				Suffix("RETURNING id").ToSql()

		if err != nil {
			return nil, err
		}

		q := db.Query{
			Name:     "order_repository.createItems",
			QueryRaw: query,
		}

		var id int
		if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
			return nil, err
		}

		itemsIDs = append(itemsIDs, id)
	}

	return itemsIDs, nil
}

func (r *repo) createOrder(ctx context.Context, order model.Order, deliveryID, paymentID int) (int, error) {
	query, args, err :=
		sq.Insert(ordersTableName).
			PlaceholderFormat(sq.Dollar).
			Columns(
				orderUIDCol,
				ordersTrackNumberCol,
				entryCol,
				deliveryIDCol,
				paymentIDCol,
				localeCol,
				internalSignatureCol,
				customerIDCol,
				deliveryServiceCol,
				shardkeyCol,
				smIDCol,
				oofShardCol,
			).
			Values(
				order.OrderUID,
				order.TrackNumber,
				order.Entry,
				deliveryID,
				paymentID,
				order.Locale,
				order.InternalSignature,
				order.CustomerID,
				order.DeliveryService,
				order.Shardkey,
				order.SmID,
				order.OofShard,
			).
			Suffix("RETURNING id").ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "order_repository.createOrder",
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) createOrdersAndItems(ctx context.Context, orderID int, itemsIDs []int) error {
	builder := sq.Insert(ordersAndItemsTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(orderIDCol, itemIDCol)

	for _, id := range itemsIDs {
		builder = builder.Values(orderID, id)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "order_repository.createOrdersAndItems",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
