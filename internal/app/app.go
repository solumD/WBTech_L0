package app

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/solumD/WBTech_L0/internal/closer"
	"github.com/solumD/WBTech_L0/internal/config"
	"github.com/solumD/WBTech_L0/internal/logger"
)

const configPath = ".env"

// App object of an app
type App struct {
	serviceProvider *serviceProvider
	server          *http.Server
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

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := a.runServer(); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	err := a.initConfig()
	if err != nil {
		return err
	}

	a.initServiceProvider()

	logger.Init(logger.GetCore(logger.GetAtomicLevel(a.serviceProvider.LoggerConfig().Level())))

	a.initServer(ctx)

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

func (a *App) initServer(ctx context.Context) {
	srv := &http.Server{
		Addr:    a.serviceProvider.ServerConfig().Address(),
		Handler: a.serviceProvider.Handler(ctx).InitRouter(),
	}

	a.server = srv
}

func (a *App) runServer() error {
	log.Printf("server is running on %s\n", a.serviceProvider.ServerConfig().Address())

	err := a.server.ListenAndServe()
	if err != nil {
		return err
	}

	log.Println("server stopped")

	return nil
}
