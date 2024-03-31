package packages

import (
	"homework/internal/order/model"
)

type Box struct{}

func (b *Box) CheckIfPackageCanBeApplied(order model.Order) bool {
	if order.Weight > 30 {
		return false
	}
	return true
}

func (b *Box) GetPrice() float64 {
	return 20
}
