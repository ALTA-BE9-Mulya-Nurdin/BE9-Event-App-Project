package response

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

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data events.Core) Event {
	return Event{

		Image:       data.Image,
		Name:        data.Name,
		Address:     data.Address,
		Date:        data.Date,
		Price:       data.Price,
		Quota:       data.Quota,
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		Link:        data.Link,
		Description: data.Description,
		Status:      data.Status,
	}
}
