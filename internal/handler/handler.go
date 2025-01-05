package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/solumD/WBTech_L0/internal/service"
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
func (h *Handler) InitRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/order/create", h.CreateOrder)
	router.Get("/order/get/{uid}", h.GetOrderByUID)

	return router
}
