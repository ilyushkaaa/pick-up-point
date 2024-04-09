package model

import "time"

type Event struct {
	RequestTime   time.Time
	RequestMethod string
	RemoteAddr    string
	URL           string
	RequestBody   string
}
