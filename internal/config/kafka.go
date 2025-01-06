package config

import (
	"errors"
	"os"
	"strings"
)

const (
	kafkaBrokersEnvName = "KAFKA_BROKERS"
)

type kafkaConfig struct {
	brokers []string
}

// NewKafkaConfig returns new kafka config
func NewKafkaConfig() (KafkaConfig, error) {
	brokersStr := os.Getenv(kafkaBrokersEnvName)
	if len(brokersStr) == 0 {
		return nil, errors.New("kafka brokers address not found")
	}

	brokers := strings.Split(brokersStr, ",")

	return &kafkaConfig{
		brokers: brokers,
	}, nil
}

// Brokers returns list of broker's addresses
func (cfg *kafkaConfig) Brokers() []string {
	if cfg.brokers == nil {
		return []string{}
	}

	return cfg.brokers
}
