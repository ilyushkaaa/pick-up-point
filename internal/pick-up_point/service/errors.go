package service

import "errors"

var (
	ErrPickUpPointNotFound      = errors.New("no pick-up points with such name")
	ErrNoPickUpPoints           = errors.New("no pick-up points in application")
	ErrPickUpPointAlreadyExists = errors.New("pick-up point with such name already exists")
)
