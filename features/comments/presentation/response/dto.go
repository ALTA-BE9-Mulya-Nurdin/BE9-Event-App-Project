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
	Event       Event
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Event struct {
	ID int `json:"id"`
}

func FromCore(data comments.Core) Comments {
	return Comments{
		ID:          data.ID,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
	}
}

func FromCoreList(data []comments.Core) []Comments {
	result := []Comments{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
