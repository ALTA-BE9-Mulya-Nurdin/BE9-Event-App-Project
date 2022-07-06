package request

import (
	"be9/event/features/comments"
)

type Comments struct {
	Description string `json:"description" form:"description"`
	EventsID    uint   `json:"events_id" form:"events_id"`
}

func ToCore(req Comments) comments.Core {
	return comments.Core{
		Description: req.Description,
		Events: comments.Events{
			ID: int(req.EventsID),
		},
	}
}
