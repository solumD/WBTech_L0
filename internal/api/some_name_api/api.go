package somenameapi

import "github.com/solumD/WBTech_L0/internal/service"

// for server or handler

// API ...
type API struct {
	orderService service.OrderService
}

// New returns new API object
func New(someService service.OrderService) *API {
	return &API{
		orderService: someService,
	}
}
