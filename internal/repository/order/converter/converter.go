package converter

import (
	modelServ "github.com/solumD/WBTech_L0/internal/model"
	modelRepo "github.com/solumD/WBTech_L0/internal/repository/order/model"
)

// FromRepoToServiceOrder converts repo order model into service order model
func FromRepoToServiceOrder(order modelRepo.Order) modelServ.Order {
	servOrder := modelServ.Order{
		OrderUID:    order.OrderUID,
		TrackNumber: order.TrackNumber,
		Entry:       order.Entry,
		Delivery: modelServ.Delivery{
			Name:    order.Delivery.Name,
			Phone:   order.Delivery.Phone,
			Zip:     order.Delivery.Zip,
			City:    order.Delivery.City,
			Address: order.Delivery.Address,
			Region:  order.Delivery.Region,
			Email:   order.Delivery.Email,
		},
		Payment: modelServ.Payment{
			Transaction:  order.Payment.Transaction,
			RequestID:    order.Payment.RequestID,
			Currency:     order.Payment.Currency,
			Provider:     order.Payment.Provider,
			Amount:       order.Payment.Amount,
			PaymentDt:    order.Payment.PaymentDt,
			Bank:         order.Payment.Bank,
			DeliveryCost: order.Payment.DeliveryCost,
			GoodsTotal:   order.Payment.GoodsTotal,
			CustomFee:    order.Payment.CustomFee,
		},
		Items:             []modelServ.Item{},
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmID:              order.SmID,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	}

	servItems := []modelServ.Item{}
	for _, item := range order.Items {
		servItem := modelServ.Item{
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			Rid:         item.Rid,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		}

		servItems = append(servItems, servItem)
	}

	servOrder.Items = servItems

	return servOrder
}
