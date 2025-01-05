package app

import (
	"context"
	"log"

	"github.com/solumD/WBTech_L0/internal/cache"
	inmemory "github.com/solumD/WBTech_L0/internal/cache/order/in_memory"
	"github.com/solumD/WBTech_L0/internal/closer"
	"github.com/solumD/WBTech_L0/internal/config"
	"github.com/solumD/WBTech_L0/internal/db"
	"github.com/solumD/WBTech_L0/internal/db/pg"
	"github.com/solumD/WBTech_L0/internal/db/transaction"
	"github.com/solumD/WBTech_L0/internal/handler"
	"github.com/solumD/WBTech_L0/internal/repository"
	orderRepo "github.com/solumD/WBTech_L0/internal/repository/order"
	"github.com/solumD/WBTech_L0/internal/service"
	orderService "github.com/solumD/WBTech_L0/internal/service/order"
)

type serviceProvider struct {
	pgConfig     config.PGConfig
	serverConfig config.ServerConfig
	loggerConfig config.LoggerConfig

	dbClient  db.Client
	txManager db.TxManager

	orderCache cache.OrderCache

	orderRepository repository.OrderRepository
	orderService    service.OrderService
	handler         *handler.Handler
}

// NewServiceProvider returns new object of service provider
func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig initializes Postgres config if it is not initialized yet and returns it
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// LoggerConfig initializes logger config if it is not initialized yet and returns it
func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %v", err)
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

// ServerConfig initializes http-server config if it is not initialized yet and returns it
func (s *serviceProvider) ServerConfig() config.ServerConfig {
	if s.serverConfig == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			log.Fatalf("failed to get http config")
		}

		s.serverConfig = cfg
	}

	return s.serverConfig
}

// DBClient initializes database client config if it is not initialized yet and returns it
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create a db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("postgres ping error: %v", err)
		}

		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager initializes transaction manager if it is not initialized yet and returns it
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// OrderCache initializes order cache if it is not initialized yet and returns it
func (s *serviceProvider) OrderCache(ctx context.Context) cache.OrderCache {
	if s.orderCache == nil {
		s.orderCache = inmemory.New()

		err := s.orderCache.LoadOrders(ctx, s.OrderRepository(ctx))
		if err != nil {
			log.Fatalf("failed to load orders in cache: %v", err)
		}

	}

	return s.orderCache
}

// OrderRepository initializes order repository if it is not initialized yet and returns it
func (s *serviceProvider) OrderRepository(ctx context.Context) repository.OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = orderRepo.New(s.DBClient(ctx))
	}

	return s.orderRepository
}

// OrderService initializes order service if it is not initialized yet and returns it
func (s *serviceProvider) OrderService(ctx context.Context) service.OrderService {
	if s.orderService == nil {
		s.orderService = orderService.New(s.OrderRepository(ctx), s.OrderCache(ctx), s.TxManager(ctx))
	}

	return s.orderService
}

// Handler initializes OrderHandler if it is not initialized yet and returns it
func (s *serviceProvider) Handler(ctx context.Context) *handler.Handler {
	if s.handler == nil {
		s.handler = handler.New(s.OrderService(ctx))
	}

	return s.handler
}
