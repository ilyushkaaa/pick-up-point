package dto

import (
	"database/sql"
	"time"

	"homework/internal/order/model"
)

type OrderDB struct {
	ID                    uint64         `db:"id"`
	ClientID              uint64         `db:"client_id"`
	Weight                float64        `db:"weight"`
	Price                 float64        `db:"price"`
	PackageType           sql.NullString `db:"package_type"`
	StorageExpirationDate time.Time      `db:"storage_expiration_date"`
	OrderIssueDate        sql.NullTime   `db:"order_issue_date"`
	IsReturned            bool           `db:"is_returned"`
}

func NewOrderDB(o model.Order) OrderDB {
	packageType := sql.NullString{Valid: false}
	if o.PackageType != "" {
		packageType.Valid = true
		packageType.String = o.PackageType
	}
	return OrderDB{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		OrderIssueDate:        sql.NullTime{Valid: false},
		PackageType:           packageType,
	}
}

func (o *OrderDB) ConvertToOrder() model.Order {
	order := model.Order{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		IsReturned:            o.IsReturned,
	}
	if o.OrderIssueDate.Valid {
		order.OrderIssueDate = o.OrderIssueDate.Time
		order.IsIssued = true
	}
	if o.PackageType.Valid {
		order.PackageType = o.PackageType.String
	}
	return order
}
