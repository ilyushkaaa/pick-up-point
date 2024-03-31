package packages

import "homework/internal/order/model"

type Wrap struct{}

func (w *Wrap) CheckIfPackageCanBeApplied(order model.Order) bool {
	return true
}

func (w *Wrap) GetPrice() float64 {
	return 1
}
