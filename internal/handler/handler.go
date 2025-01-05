package handler

import "github.com/solumD/WBTech_L0/internal/service"

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
