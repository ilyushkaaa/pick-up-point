package dto

import (
	"time"

	modelOrder "homework/internal/order/model"
)

type OrderOutput struct {
	ID                    uint64    `json:"id"`
	ClientID              uint64    `json:"client_id"`
	Weight                float64   `json:"weight"`
	Price                 float64   `json:"price"`
	PackageType           string    `json:"package_type,omitempty"`
	StorageExpirationDate time.Time `json:"storage_expiration_date"`
	OrderIssueDate        time.Time `json:"order_issue_date,omitempty"`
	IsIssued              bool      `json:"is_issued"`
	IsReturned            bool      `json:"is_returned"`
	PickUpPointID         uint64    `json:"pick_up_point_id"`
}

func NewOrderOutput(o modelOrder.Order) OrderOutput {
	return OrderOutput{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		PackageType:           o.PackageType,
		StorageExpirationDate: o.StorageExpirationDate,
		OrderIssueDate:        o.OrderIssueDate,
		IsIssued:              o.IsIssued,
		IsReturned:            o.IsReturned,
		PickUpPointID:         o.PickUpPointID,
	}
}
