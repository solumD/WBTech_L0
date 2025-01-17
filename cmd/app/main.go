package main

import (
	"context"
	"log"

	"github.com/solumD/WBTech_L0/internal/app"
)

// @title WBTech_L0 order service
// @version 1.0
// @host localhost:8080
// @basePath /

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
