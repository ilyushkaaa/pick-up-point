package dto

import "github.com/asaskevich/govalidator"

type ReturnOrderInputData struct {
	OrderID  uint64 `json:"order_id" valid:"required"`
	ClientID uint64 `json:"client_id" valid:"required"`
}

func (p *ReturnOrderInputData) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
