package packages

import "homework/internal/order/model"

type Package interface {
	CheckIfPackageCanBeApplied(order model.Order) bool
	GetPrice() float64
}
