package presentation

import "be9/event/features/eventdetail"

type EventDetailHandler struct {
	eventDetailBusiness eventdetail.Business
}

func NewEventDetailHandler(business eventdetail.Business) *EventDetailHandler {
	return &EventDetailHandler{
		eventDetailBusiness: business,
	}
}
