package dto

import (
	"database/sql"
	"time"

	"homework/internal/order/model"
)

type OrderDB struct {
	ID                    uint64       `db:"id"`
	ClientID              uint64       `db:"client_id"`
	Weight                float64      `db:"weight"`
	Price                 float64      `db:"price"`
	StorageExpirationDate time.Time    `db:"storage_expiration_date"`
	OrderIssueDate        sql.NullTime `db:"order_issue_date"`
	IsIssued              bool         `db:"is_issued"`
	IsReturned            bool         `db:"is_returned"`
}

func NewOrderDB(o model.Order) OrderDB {
	return OrderDB{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		OrderIssueDate:        sql.NullTime{Valid: false},
	}
}

func (o *OrderDB) ConvertToOrder() model.Order {
	order := model.Order{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		IsIssued:              o.IsIssued,
		IsReturned:            o.IsReturned,
	}
	if o.IsIssued && o.OrderIssueDate.Valid {
		order.OrderIssueDate = o.OrderIssueDate.Time
	}
	return order
}
