package data

import (
	"be9/event/features/categorys"

	"gorm.io/gorm"
)

type Categorys struct {
	gorm.Model
	CategoryName string `json:"category_name" form:"category_name"`
}

func (data *Categorys) toCore() categorys.Core {
	return categorys.Core{
		ID:           int(data.ID),
		CategoryName: data.CategoryName,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func toCoreList(data []Categorys) []categorys.Core {
	result := []categorys.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core categorys.Core) Categorys {
	return Categorys{
		CategoryName: core.CategoryName,
	}
}
