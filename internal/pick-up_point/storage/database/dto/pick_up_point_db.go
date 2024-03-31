package dto

type PickUpPointDB struct {
	ID          uint64 `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Region      string `db:"region"`
	City        string `db:"city"`
	Street      string `db:"street"`
	HouseNum    string `db:"house_num"`
}
