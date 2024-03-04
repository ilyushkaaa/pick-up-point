package service

import "errors"

var (
	ErrOrderAlreadyExists       = errors.New("order with such ID already exists")
	ErrShelfTimeExpired         = errors.New("shelf time for order has already expired")
	ErrOrderAlreadyIssued       = errors.New("order has been already issued")
	ErrOrderShelfLifeNotExpired = errors.New("order shelf life has not been expired")
	ErrOrderNotFound            = errors.New("order was not found")
	ErrOrdersOfDifferentClients = errors.New("orders belong to different clients")
	ErrClientOrderNotFound      = errors.New("user has not got orders with such ID")
	ErrReturnTimeExpired        = errors.New("time of order return gas expired")
	ErrOrderIsNotIssued         = errors.New("order is not issued")
	ErrOrderIsReturned          = errors.New("order is returned")
	ErrNoOrdersOnThisPage       = errors.New("no orders on page with this number")
)
