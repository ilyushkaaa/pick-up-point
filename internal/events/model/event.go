package model

import "time"

type Event struct {
	RequestTime   string
	RequestMethod string
	RemoteAddr    string
	URL           string
}

func NewEvent(rAddr, rURL, rMethod string) Event {
	return Event{
		RequestTime:   time.Now().Format("2006-01-02 15:04:05"),
		RequestMethod: rMethod,
		RemoteAddr:    rAddr,
		URL:           rURL,
	}
}
