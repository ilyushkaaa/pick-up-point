package fixtures

import (
	"homework/internal/events/model"
	"homework/tests/states"
)

type EventBuilder struct {
	instance *model.Event
}

func Event() *EventBuilder {
	return &EventBuilder{instance: &model.Event{}}
}

func (b *EventBuilder) RequestTime(v string) *EventBuilder {
	b.instance.RequestTime = v
	return b
}

func (b *EventBuilder) RequestMethod(v string) *EventBuilder {
	b.instance.RequestMethod = v
	return b
}

func (b *EventBuilder) RemoteAddr(v string) *EventBuilder {
	b.instance.RemoteAddr = v
	return b
}

func (b *EventBuilder) Pointer() *model.Event {
	return b.instance
}

func (b *EventBuilder) Value() model.Event {
	return *b.instance
}

func (b *EventBuilder) Valid() *EventBuilder {
	return Event().RequestTime(states.EventTime).RequestMethod(states.EventMethodGet).RemoteAddr(states.EventAddrLocal)
}
