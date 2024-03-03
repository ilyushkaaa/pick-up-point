package myerrors

import "errors"

var (
	ErrorOrderAlreadyExists       = errors.New("order with such ID already exists")
	ErrorShelfTimeExpired         = errors.New("shelf time for order has already expired")
	ErrorOrderAlreadyIssued       = errors.New("order has been already issued")
	ErrorOrderShelfLifeNotExpired = errors.New("order shelf life has not been expired")
	ErrorOrderNotFound            = errors.New("order was not found")
	ErrorOrdersOfDifferentClients = errors.New("orders belong to different clients")
	ErrorClientOrderNotFound      = errors.New("user has not got orders with such ID")
	ErrorReturnTimeExpired        = errors.New("time of order return gas expired")
	ErrorOrderIsNotIssued         = errors.New("order is not issued")
	ErrorOrderIsReturned          = errors.New("order is returned")
)
