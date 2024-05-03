package dto

type AddressDTO struct {
	Region   string `json:"region" valid:"required,length(3|50)"`
	City     string `json:"city" valid:"required,length(3|50)"`
	Street   string `json:"street" valid:"required,length(2|50)"`
	HouseNum string `json:"house_num" valid:"required,length(1|10)"`
}
