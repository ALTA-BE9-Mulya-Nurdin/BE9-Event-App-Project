package request

import (
	"be9/event/features/events"
	"be9/event/features/participants"
)

type Participants struct {
	EventsID uint `json:"events_id" form:"events_id"`
}

func ToCore(req Participants) participants.Core {
	return participants.Core{
		Events: events.Core{
			ID: int(req.EventsID),
		},
	}
}
