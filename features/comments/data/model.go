package data

import (
	"be9/event/features/comments"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Description string `json:"description" form:"description"`
	// User User
	// Event Event
}

// type User struct {
// 	ID   int
// }

// type Event struct {
// 	ID   int
// }

func (data *Comments) toCore() comments.Core {
	return comments.Core{
		ID:          int(data.ID),
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func toCoreList(data []Comments) []comments.Core {
	result := []comments.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core comments.Core) Comments {
	return Comments{
		Description: core.Description,
	}
}
