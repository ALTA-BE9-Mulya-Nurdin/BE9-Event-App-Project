package business

import (
	"be9/event/features/events"
	"be9/event/features/users"
)

type eventUsecase struct {
	eventData events.Data
}

func NewEventBusiness(evntData events.Data) users.Business {
	return &eventUsecase{
		eventData: evntData,
	}
}
