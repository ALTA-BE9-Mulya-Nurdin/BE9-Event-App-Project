package request

import (
	"be9/event/features/events"
	"time"
)

type Event struct {
	IDUsers     int       `json:"id_users" from:"id_users"`
	IDCategory  int       `json:"id_category" from:"id_category"`
	Image       string    `json:"image" from:"image"`
	Name        string    `json:"name" from:"name"`
	Address     string    `json:"address" from:"address"`
	Date        time.Time `json:"date" from:"date"`
	Price       int       `json:"price" from:"price"`
	Quota       int       `json:"quota" from:"quota"`
	Longitude   string    `json:"longitude" from:"longitude"`
	Latitude    string    `json:"latitude" from:"latitude"`
	Link        string    `json:"link" from:"link"`
	Description string    `json:"description" from:"description"`
	Status      string    `json:"status" from:"status"`
}

func ToCore(req Event) events.Core {
	return events.Core{
		Image:       req.Image,
		Name:        req.Name,
		Address:     req.Address,
		Date:        req.Date,
		Price:       req.Price,
		Quota:       req.Quota,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Link:        req.Link,
		Description: req.Description,
		Status:      req.Status,
	}
}
