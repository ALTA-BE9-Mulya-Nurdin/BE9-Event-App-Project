package request

import "be9/event/features/events"

type Events struct {
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	Date        string `json:"date" form:"date"`
	Price       int    `json:"price" form:"price"`
	Quote       int    `json:"quote" form:"quote"`
	Longitude   string `json:"longitude" form:"longitude"`
	Latitude    string `json:"latitude" form:"latitude"`
	Link        string `json:"link" form:"link"`
	Description string `json:"description" form:"description"`
	Status      string `json:"status" form:"status"`
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
	}
}
