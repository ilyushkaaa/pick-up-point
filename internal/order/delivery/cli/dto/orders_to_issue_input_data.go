package dto

import "github.com/asaskevich/govalidator"

type OrdersToIssueInputData struct {
	OrdersIDs []uint64 `json:"orders-ids" valid:"required"`
}

func (p *OrdersToIssueInputData) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
