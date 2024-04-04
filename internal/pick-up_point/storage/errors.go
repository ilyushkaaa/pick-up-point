package storage

import "errors"

var (
	ErrPickUpPointNotFound   = errors.New("no pick-up points with such id")
	ErrPickUpPointNameExists = errors.New("pick-up point with such name already exists")
)
