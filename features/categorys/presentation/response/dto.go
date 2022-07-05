package response

import (
	"time"

	"be9/event/features/categorys"
)

type Categorys struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func FromCore(data categorys.Core) Categorys {
	return Categorys{
		ID:           data.ID,
		CategoryName: data.CategoryName,
		CreatedAt:    data.CreatedAt,
	}
}

func FromCoreList(data []categorys.Core) []Categorys {
	result := []Categorys{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
