package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/solumD/WBTech_L0/internal/logger"
	"github.com/solumD/WBTech_L0/internal/model"
	"go.uber.org/zap"
)

// Consume gets orders from msgIn chan and saves them in repository
func (s *srv) Consume(ctx context.Context, msgIn chan *sarama.ConsumerMessage) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-msgIn:
			if !ok {
				return fmt.Errorf("message chan is closed")
			}

			order := &model.Order{}
			if err := json.Unmarshal(msg.Value, order); err != nil {
				return err
			}

			logger.Info("recieved order from consumer", zap.Any("order", *order))

			if err := s.CreateOrder(ctx, *order); err != nil {
				return err
			}
		}

	}
}
