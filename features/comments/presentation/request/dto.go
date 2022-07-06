package request

import "be9/event/features/comments"

type Comments struct {
	Description string `json:"description" form:"description"`
	EventID     uint   `json:"event_id" form:"event_id"`
}

func ToCore(req Comments) comments.Core {
	return comments.Core{
		Description: req.Description,
		Event: comments.Event{
			ID: int(req.EventID),
		},
	}
}
