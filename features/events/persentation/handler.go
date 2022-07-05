package persentation

import "be9/event/features/events"

type EventHandler struct {
	eventBisnis events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		eventBisnis: business,
	}
}
