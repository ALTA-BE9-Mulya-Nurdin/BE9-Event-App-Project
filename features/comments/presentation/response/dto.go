package response

import (
	"be9/event/features/comments"
	"time"
)

type Comments struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	User        User
	Event       Events
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Events struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data comments.Core) Comments {
	return Comments{
		ID:          data.ID,
		Description: data.Description,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
		},
		Event: Events{
			ID:   data.Events.ID,
			Name: data.Events.Name,
		},
	}
}

func FromCoreList(data []comments.Core) []Comments {
	result := []Comments{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
