package validate

import (
	"fmt"
	"strings"

	"github.com/solumD/WBTech_L0/internal/model"
)

// Order validates order
func Order(order model.Order) error {
	if len(order.OrderUID) == 0 || strings.Contains(order.OrderUID, " ") {
		return fmt.Errorf("invalid order uid")
	}

	if len(order.TrackNumber) == 0 || strings.Contains(order.TrackNumber, " ") {
		return fmt.Errorf("invalid order track number")
	}

	if len(order.Entry) == 0 || strings.Contains(order.Entry, " ") {
		return fmt.Errorf("invalid order entry")
	}

	if len(order.Locale) == 0 || strings.Contains(order.Locale, " ") {
		return fmt.Errorf("invalid order locale")
	}

	if len(order.InternalSignature) != 0 && strings.Contains(order.InternalSignature, " ") {
		return fmt.Errorf("invalid order internal signature")
	}

	if len(order.CustomerID) == 0 || strings.Contains(order.CustomerID, " ") {
		return fmt.Errorf("invalid order customer id")
	}

	if len(order.DeliveryService) == 0 {
		return fmt.Errorf("invalid order delivery service")
	}

	if len(order.Shardkey) == 0 || strings.Contains(order.Shardkey, " ") {
		return fmt.Errorf("invalid order shard key")
	}

	if order.SmID <= 0 {
		return fmt.Errorf("invalid order sm id")
	}

	if len(order.OofShard) == 0 || strings.Contains(order.OofShard, " ") {
		return fmt.Errorf("invalid oof shard key")
	}

	if err := Delivery(order.Delivery); err != nil {
		return err
	}

	if err := Payment(order.Payment); err != nil {
		return err
	}

	if err := Items(order.Items); err != nil {
		return err
	}

	return nil
}

// Delivery validates delivery
func Delivery(delivery model.Delivery) error {
	if len(delivery.Name) == 0 {
		return fmt.Errorf("invalid delivery name")
	}

	if len(delivery.Phone) == 0 {
		return fmt.Errorf("invalid delivery phone")
	}

	if len(delivery.Zip) == 0 || strings.Contains(delivery.Zip, " ") {
		return fmt.Errorf("invalid delivery zip")
	}

	if len(delivery.City) == 0 {
		return fmt.Errorf("invalid delivery city")
	}

	if len(delivery.Address) == 0 {
		return fmt.Errorf("invalid delivery address")
	}

	if len(delivery.Region) == 0 {
		return fmt.Errorf("invalid delivery region")
	}

	if len(delivery.Email) == 0 || strings.Contains(delivery.Email, " ") {
		return fmt.Errorf("invalid delivery email")
	}

	return nil
}

// Payment validates payment
func Payment(payment model.Payment) error {
	if len(payment.Transaction) == 0 || strings.Contains(payment.Transaction, " ") {
		return fmt.Errorf("invalid payment transaction")
	}

	if len(payment.RequestID) != 0 && strings.Contains(payment.RequestID, " ") {
		return fmt.Errorf("invalid payment request id")
	}

	if len(payment.Currency) == 0 || strings.Contains(payment.Currency, " ") {
		return fmt.Errorf("invalid payment currency")
	}

	if len(payment.Provider) == 0 {
		return fmt.Errorf("invalid payment provider")
	}

	if payment.Amount <= 0 {
		return fmt.Errorf("invalid payment amount")
	}

	if payment.PaymentDt <= 0 {
		return fmt.Errorf("invalid payment payment dt")
	}

	if len(payment.Bank) == 0 || strings.Contains(payment.Bank, " ") {
		return fmt.Errorf("invalid payment bank")
	}

	if payment.DeliveryCost <= 0 {
		return fmt.Errorf("invalid payment delivery cost")
	}

	if payment.GoodsTotal <= 0 {
		return fmt.Errorf("invalid payment goods total")
	}

	if payment.CustomFee <= 0 {
		return fmt.Errorf("invalid payment custom fee")
	}

	return nil
}

// Items validates items
func Items(items []model.Item) error {
	if len(items) == 0 {
		return fmt.Errorf("invalid items: items amount can't be nil")
	}

	for idx, item := range items {
		if item.ChrtID <= 0 {
			return fmt.Errorf("item idx %d: invalid chrt id", idx)
		}

		if len(item.TrackNumber) == 0 || strings.Contains(item.TrackNumber, " ") {
			return fmt.Errorf("item idx %d: invalid track nubmer", idx)
		}

		if item.Price <= 0 {
			return fmt.Errorf("item idx %d: invalid price", idx)
		}

		if len(item.Rid) == 0 || strings.Contains(item.Rid, " ") {
			return fmt.Errorf("item idx %d: invalid rid", idx)
		}

		if len(item.Name) == 0 {
			return fmt.Errorf("item idx %d: invalid name", idx)
		}

		if item.Sale <= 0 || item.Sale > 100 {
			return fmt.Errorf("item idx %d: invalid sale", idx)
		}

		if len(item.Size) == 0 || strings.Contains(item.Size, " ") {
			return fmt.Errorf("item idx %d: invalid size", idx)
		}

		if item.TotalPrice <= 0 {
			return fmt.Errorf("item idx %d: invalid total price", idx)
		}

		if item.NmID <= 0 {
			return fmt.Errorf("item idx %d: invalid nm id", idx)
		}

		if len(item.Brand) == 0 {
			return fmt.Errorf("item idx %d: invalid brand", idx)
		}

		if item.Status <= 0 {
			return fmt.Errorf("item idx %d: invalid status", idx)
		}
	}

	return nil
}
