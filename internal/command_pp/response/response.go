package response

import "github.com/google/uuid"

type Response struct {
	ID   uuid.UUID
	Body string
	Err  error
}
