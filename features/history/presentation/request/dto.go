package request

import (
	"be9/event/features/events"
	"be9/event/features/participants"
)

// type Historys struct {
// 	// UsersID  uint `json:"users_id" form:"users_id"`
// 	EventsID uint `json:"events_id" form:"events_id"`
// }

// func ToCore(req Historys) history.Core {
// 	return history.Core{
// 		Users: users.Core{
// 			ID: int(req.UsersID),
// 		},
// 		Events: events.Core{
// 			ID: int(req.EventsID),
// 		},
// 	}
// }

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
