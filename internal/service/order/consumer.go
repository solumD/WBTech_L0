package order

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/IBM/sarama"
	"github.com/solumD/WBTech_L0/internal/logger"
	"github.com/solumD/WBTech_L0/internal/model"
	"go.uber.org/zap"
)

// Consume gets orders from msgIn chan and saves them in repository
func (s *srv) Consume(ctx context.Context, msgIn chan *sarama.ConsumerMessage) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-msgIn:
				if !ok {
					logger.Info("message chan is closed, stopping service consumer")
					return
				}

				order := &model.Order{}
				if err := json.Unmarshal(msg.Value, order); err != nil {
					logger.Error("failed to unmarshal order", zap.Error(err))

				}

				logger.Info("recieved order from consumer", zap.Any("order", *order))

				if err := s.CreateOrder(ctx, *order); err != nil {
					logger.Error("failed to create order", zap.Error(err))
				}
			}

		}
	}()

	wg.Wait()
}
