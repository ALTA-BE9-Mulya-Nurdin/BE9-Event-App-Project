package data

import (
	"be9/event/features/events"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Image       string    `json:"image" form:"image"`
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

func toCoreList(data []Event) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:          int(data.ID),
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
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func fromCore(core events.Core) Event {
	return Event{
		Image:       core.Image,
		Name:        core.Name,
		Address:     core.Address,
		Date:        core.Date,
		Price:       core.Price,
		Quota:       core.Quota,
		Longitude:   core.Longitude,
		Latitude:    core.Latitude,
		Link:        core.Link,
		Description: core.Description,
		Status:      core.Status,
	}
}
