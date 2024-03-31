package response

import "errors"

var (
	ErrInternal    = errors.New("internal error")
	ErrInvalidJSON = errors.New("invalid json passed")
)
