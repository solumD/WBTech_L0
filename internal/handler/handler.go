package handler

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/solumD/WBTech_L0/docs"
	"github.com/solumD/WBTech_L0/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Handler object, which contains all handlers
type Handler struct {
	orderService service.OrderService
}

// New returns new handler object
func New(orderService service.OrderService) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

// InitRouter inits router's routes and returns router
func (h *Handler) InitRouter(serverAddress string) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/order/{uid}", h.GetOrderByUID)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", serverAddress)),
	))

	return router
}
