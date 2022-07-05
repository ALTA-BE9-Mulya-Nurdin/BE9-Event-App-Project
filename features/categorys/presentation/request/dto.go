package request

import (
	"be9/event/features/categorys"
)

type Categorys struct {
	CategoryName string `json:"category_name" form:"category_name"`
}

func ToCore(req Categorys) categorys.Core {
	return categorys.Core{
		CategoryName: req.CategoryName,
	}
}
