package converter

import (
	modelServ "github.com/solumD/WBTech_L0/internal/model"
	modelRepo "github.com/solumD/WBTech_L0/internal/repository/order/model"
)

// FromRepoToServiceOrder converts repo order model into service order model
func FromRepoToServiceOrder(_ *modelRepo.Order) *modelServ.Order {
	return nil
}
