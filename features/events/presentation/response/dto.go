package response

import (
	"be9/event/features/events"
	"time"
)

type Events struct {
	ID          int       `json:"id"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Date        string    `json:"date"`
	Price       int       `json:"price"`
	Quote       int       `json:"quote"`
	Longitude   string    `json:"longitude"`
	Latitude    string    `json:"latitude"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromCoreList(data []events.Core) []Events {
	result := []Events{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data events.Core) Events {
	return Events{
		ID:          data.ID,
		Image:       data.Image,
		Name:        data.Name,
		Address:     data.Address,
		Date:        data.Date,
		Price:       data.Price,
		Quote:       data.Quote,
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		Link:        data.Link,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}
}
