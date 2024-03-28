package packages

import (
	"homework/internal/order/model"
)

type Packet struct{}

func (p *Packet) AddPackageToOrder(order *model.Order) error {
	if order.Weight > 10 {
		return ErrPackageCanNotBeApplied
	}
	order.Price += 5
	return nil
}
