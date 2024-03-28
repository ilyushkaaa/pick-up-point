package packages

import (
	"homework/internal/order/model"
)

type Box struct{}

func (b *Box) AddPackageToOrder(order *model.Order) error {
	if order.Weight > 30 {
		return ErrPackageCanNotBeApplied
	}
	order.Price += 20
	return nil
}
