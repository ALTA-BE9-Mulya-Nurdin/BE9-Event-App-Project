package request

import (
	"be9/event/features/eventdetail"
	"be9/event/features/events"
)

type EventDetail struct {
	EventsID uint `json:"events_id" form:"events_id"`
}

func ToCore(req EventDetail) eventdetail.Core {
	return eventdetail.Core{
		Events: events.Core{
			ID: int(req.EventsID),
		},
	}
}
