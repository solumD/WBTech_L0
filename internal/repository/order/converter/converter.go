package converter

import (
	modelServ "github.com/solumD/WBTech_L0/internal/model"
	modelRepo "github.com/solumD/WBTech_L0/internal/repository/order/model"
)

// FromRepoToServiceOrder gets models from repo, connects them and converts into one service order model
func FromRepoToServiceOrder(order *modelRepo.Order) *modelServ.Order {
	return nil
}
