package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/solumD/WBTech_L0/internal/model"
)

// Simple producer example for demonstration

const (
	brokerAddress = "localhost:9092"
	topicName     = "orders-topic"
)

func main() {
	producer, err := newSyncProducer([]string{brokerAddress})
	if err != nil {
		log.Fatalf("failed to start producer: %v\n", err.Error())
	}

	defer func() {
		if err = producer.Close(); err != nil {
			log.Fatalf("failed to close producer: %v\n", err.Error())
		}
	}()

	for {
		order := generateOrder()
		data, err := json.Marshal(order)
		if err != nil {
			log.Printf("failed to marshal order: %v\n", err)
		}

		msg := &sarama.ProducerMessage{
			Topic: topicName,
			Value: sarama.ByteEncoder(data),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("failed to send message in Kafka: %v\n", err.Error())
			return
		}

		log.Printf("message sent to partition %d with offset %d\n", partition, offset)
		time.Sleep(5 * time.Second)
	}
}

func newSyncProducer(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func generateOrder() model.Order {
	return model.Order{
		OrderUID:          gofakeit.UUID(),
		TrackNumber:       gofakeit.UUID(),
		Entry:             gofakeit.Word(),
		Delivery:          generateDelivery(),
		Payment:           generatePayment(),
		Items:             generateItems(),
		Locale:            gofakeit.LanguageAbbreviation(),
		InternalSignature: gofakeit.Word(),
		CustomerID:        gofakeit.Word(),
		DeliveryService:   gofakeit.Word(),
		Shardkey:          gofakeit.Word(),
		SmID:              gofakeit.Number(1, 1000),
		OofShard:          gofakeit.Word(),
	}
}

func generateDelivery() model.Delivery {
	return model.Delivery{
		Name:    gofakeit.Name(),
		Phone:   gofakeit.Phone(),
		Zip:     gofakeit.Zip(),
		City:    gofakeit.City(),
		Address: gofakeit.Street(),
		Region:  gofakeit.State(),
		Email:   gofakeit.Email(),
	}
}

func generatePayment() model.Payment {
	return model.Payment{
		Transaction:  gofakeit.UUID(),
		RequestID:    gofakeit.UUID(),
		Currency:     gofakeit.CurrencyShort(),
		Provider:     gofakeit.Word(),
		Amount:       gofakeit.Number(1, 10000),
		PaymentDt:    gofakeit.Number(1, 100000000),
		Bank:         gofakeit.Word(),
		DeliveryCost: gofakeit.Number(100, 10000),
		GoodsTotal:   gofakeit.Number(1, 1000),
		CustomFee:    gofakeit.Number(0, 100),
	}
}

func generateItems() []model.Item {
	count := gofakeit.Number(1, 5)
	items := make([]model.Item, count)
	for i := 0; i < count; i++ {
		items[i] = model.Item{
			ChrtID:      gofakeit.Number(1, 100000),
			TrackNumber: gofakeit.UUID(),
			Price:       gofakeit.Number(1, 10000),
			Rid:         gofakeit.UUID(),
			Name:        gofakeit.Word(),
			Sale:        gofakeit.Number(0, 100),
			Size:        gofakeit.Word(),
			TotalPrice:  gofakeit.Number(1, 10000),
			NmID:        gofakeit.Number(1, 10000),
			Brand:       gofakeit.Company(),
			Status:      gofakeit.HTTPStatusCode(),
		}
	}
	return items
}
