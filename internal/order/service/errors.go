package service

import "errors"

var (
	ErrShelfTimeExpired          = errors.New("shelf time for order has already expired")
	ErrOrderAlreadyIssued        = errors.New("order has been already issued")
	ErrOrderShelfLifeNotExpired  = errors.New("order shelf life has not been expired")
	ErrNotAllOrdersWereFound     = errors.New("some of orders were not found")
	ErrOrdersOfDifferentClients  = errors.New("orders belong to different clients")
	ErrClientOrderNotFound       = errors.New("user has not got orders with such ID")
	ErrReturnTimeExpired         = errors.New("time of order return gas expired")
	ErrOrderIsNotIssued          = errors.New("order is not issued")
	ErrOrderIsReturned           = errors.New("order is returned")
	ErrNoOrdersOnThisPage        = errors.New("no orders on page with this number")
	ErrOrderAlreadyInPickUpPoint = errors.New("order with such id is already in pick-up point")
	ErrUnknownPackage            = errors.New("unknown package type passed")
)
