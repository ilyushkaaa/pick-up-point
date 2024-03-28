package packages

import "homework/internal/order/model"

type Package interface {
	AddPackageToOrder(order *model.Order) error
}
