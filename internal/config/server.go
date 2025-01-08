package config

import (
	"errors"
	"net"
	"os"
	"time"
)

const (
	serverHostEnvName = "SERVER_HOST"
	serverPortEnvName = "SERVER_PORT"
	serverTimeout     = 5
	serverIdleTimeout = 30
)

type serverConfig struct {
	host        string
	port        string
	timeout     time.Duration
	idleTimeout time.Duration
}

// NewServerConfig returns new server config
func NewServerConfig() (ServerConfig, error) {
	host := os.Getenv(serverHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("server host not found")
	}

	port := os.Getenv(serverPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("server port not found")
	}

	return &serverConfig{
		host:        host,
		port:        port,
		timeout:     time.Second * serverTimeout,
		idleTimeout: time.Second * serverIdleTimeout,
	}, nil
}

// Address returns full address of a server
func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

// Timeout returns timeout of a server
func (cfg *serverConfig) Timeout() time.Duration {
	return cfg.timeout
}

// IdleTimeout returns idle timeout of a server
func (cfg *serverConfig) IdleTimeout() time.Duration {
	return cfg.idleTimeout
}
