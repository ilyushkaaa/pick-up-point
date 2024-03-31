package storage

import "errors"

var (
	ErrOrderAlreadyExists  = errors.New("order with such ID already exists")
	ErrOrderNotFound       = errors.New("order was not found")
	ErrClientOrderNotFound = errors.New("user has not got orders with such ID")
)
