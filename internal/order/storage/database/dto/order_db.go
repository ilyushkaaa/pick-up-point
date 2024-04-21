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
	PickUpPointID         uint64         `db:"pick_up_point_id"`
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
		PickUpPointID:         o.PickUpPointID,
	}
}
