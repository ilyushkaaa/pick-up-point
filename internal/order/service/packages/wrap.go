package packages

import "homework/internal/order/model"

type Wrap struct{}

func (w *Wrap) AddPackageToOrder(order *model.Order) error {
	order.Price += 1
	return nil
}
