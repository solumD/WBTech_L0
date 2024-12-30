package app

import (
	"context"

	"github.com/solumD/WBTech_L0/internal/closer"
	"github.com/solumD/WBTech_L0/internal/config"
	"github.com/solumD/WBTech_L0/internal/logger"
)

const configPath = ".env"

// App object of an app
type App struct {
	serviceProvider *serviceProvider
	// servers or handlers
}

// NewApp returns new App object
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts an App
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	/*testOrder := model.Order{
		OrderUID:    "new-order-uid",
		TrackNumber: "different-track-number",
		Entry:       "Alternative Entry",
		Delivery: model.Delivery{
			Name:    "Jane Smith",
			Phone:   "+0987654321",
			Zip:     "54321",
			City:    "New City",
			Address: "Different Address",
			Region:  "Another Region",
			Email:   "janesmith@example.com",
		},
		Payment: model.Payment{
			Transaction:  "txn_9876",
			RequestID:    "req_4321",
			Currency:     "EUR",
			Provider:     "Other Provider",
			Amount:       200,
			PaymentDt:    1667890123,
			Bank:         "Other Bank",
			DeliveryCost: 15,
			GoodsTotal:   185,
			CustomFee:    7,
		},
		Items: []model.Item{
			{
				ChrtID:      11,
				TrackNumber: "alt-item-track-num-1",
				Price:       75,
				Rid:         "alt-item-rid-1",
				Name:        "Alternate Item 1",
				Sale:        25,
				Size:        "S",
				TotalPrice:  60,
				NmID:        201,
				Brand:       "Alt Brand",
				Status:      3,
			},
			{
				ChrtID:      22,
				TrackNumber: "alt-item-track-num-2",
				Price:       110,
				Rid:         "alt-item-rid-2",
				Name:        "Alternate Item 2",
				Sale:        35,
				Size:        "XL",
				TotalPrice:  80,
				NmID:        202,
				Brand:       "Yet Another Alt Brand",
				Status:      4,
			},
		},
		Locale:            "ru-RU",
		InternalSignature: "alternate-signature",
		CustomerID:        "cust_456",
		DeliveryService:   "Alternative Delivery Service",
		Shardkey:          "alternative-shard-key",
		SmID:              8888,
		DateCreated:       time.Now().AddDate(0, 0, -14), // Two weeks ago
		OofShard:          "alternate-oof-shard",
	}

	err := a.serviceProvider.OrderRepository(context.Background()).CreateOrder(context.TODO(), testOrder)
	if err != nil {
		return err
	}*/

	_, err := a.serviceProvider.OrderRepository(context.Background()).GetAllOrders(context.TODO())
	if err != nil {
		return err
	}
	// some gorutines with running servers
	return nil
}

func (a *App) initDeps(_ context.Context) error {
	err := a.initConfig()
	if err != nil {
		return err
	}

	a.initServiceProvider()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(a.serviceProvider.LoggerConfig().Level())))

	return nil
}

func (a *App) initConfig() error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider() {
	a.serviceProvider = NewServiceProvider()
}
