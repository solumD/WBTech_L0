package config

import (
	"time"

	"github.com/joho/godotenv"
)

// PGConfig interface of Postgres config
type PGConfig interface {
	DSN() string
}

// KafkaConfig interface ok Kafka config
type KafkaConfig interface {
	Brokers() []string
}

// LoggerConfig interface of logger config
type LoggerConfig interface {
	Level() string
}

// ServerConfig interface of Server config
type ServerConfig interface {
	Address() string
	Timeout() time.Duration
	IdleTimeout() time.Duration
}

// Load reads the .env file at the specified path
// and loads the variables into the project
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
