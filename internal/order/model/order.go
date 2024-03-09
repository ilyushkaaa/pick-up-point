package model

import "time"

type Order struct {
	ID                    int
	ClientID              int
	StorageExpirationDate time.Time
	OrderIssueDate        time.Time
	IsIssued              bool
	IsReturned            bool
}
