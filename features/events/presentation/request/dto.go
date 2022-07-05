package request

import (
	"be9/event/features/categorys"
	"be9/event/features/events"
)

type Events struct {
	Image       string `json:"image" validate:"required" form:"image"`
	Name        string `json:"name" validate:"required" form:"name"`
	Address     string `json:"address" validate:"required" form:"address"`
	Date        string `json:"date" validate:"required" form:"date"`
	Price       int    `json:"price" validate:"required" form:"price"`
	Quote       int    `json:"quote" validate:"required" form:"quote"`
	Longitude   string `json:"longitude" form:"longitude"`
	Latitude    string `json:"latitude" form:"latitude"`
	Link        string `json:"link" form:"link"`
	Description string `json:"description" validate:"required" form:"description"`
	Status      string `json:"status" validate:"required" form:"status"`
	CategorysID uint   `json:"categorys_id" validate:"required" form:"categorys_id"`
}

func ToCore(req Events) events.Core {
	return events.Core{
		Image:       req.Image,
		Name:        req.Name,
		Address:     req.Address,
		Date:        req.Date,
		Price:       req.Price,
		Quote:       req.Quote,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Link:        req.Link,
		Description: req.Description,
		Status:      req.Status,
		Categorys: categorys.Core{
			ID: int(req.CategorysID),
		},
	}
}
