package packages

import (
	"homework/internal/order/model"
)

type Packet struct{}

func (p *Packet) CheckIfPackageCanBeApplied(order model.Order) bool {
	if order.Weight > 10 {
		return false
	}
	return true
}

func (p *Packet) GetPrice() float64 {
	return 5
}
