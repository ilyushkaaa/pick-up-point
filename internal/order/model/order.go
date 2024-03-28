package model

import "time"

type Order struct {
	ID                    uint64
	ClientID              uint64
	Weight                float64
	Price                 float64
	StorageExpirationDate time.Time
	OrderIssueDate        time.Time
	IsIssued              bool
	IsReturned            bool
}

func NewOrder(id, clientID uint64, weight, price float64, expireDate time.Time) Order {
	return Order{
		ID:                    id,
		ClientID:              clientID,
		Weight:                weight,
		Price:                 price,
		StorageExpirationDate: expireDate,
	}
}
