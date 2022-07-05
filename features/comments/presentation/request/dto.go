package request

import "be9/event/features/comments"

type Comments struct {
	Description string `json:"description" form:"description"`
	// User User
	// Event Event
}

func ToCore(req Comments) comments.Core {
	return comments.Core{
		Description: req.Description,
	}
}
