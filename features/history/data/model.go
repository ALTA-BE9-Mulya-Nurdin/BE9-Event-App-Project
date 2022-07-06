package data

import (
	"be9/event/features/events"
	_events "be9/event/features/events/data"
	"be9/event/features/history"
	"be9/event/features/users"
	_users "be9/event/features/users/data"

	"gorm.io/gorm"
)

type Participants struct {
	gorm.Model
	UserID   uint `json:"user_id" form:"user_id"`
	EventsID uint `json:"events_id" form:"events_id"`
	Events   _events.Events
	User     _users.User
}

func toCoreList(data []Participants) []history.Core {
	result := []history.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Participants) toCore() history.Core {
	return history.Core{
		ID: int(data.ID),
		Events: events.Core{
			ID:          int(data.Events.ID),
			Image:       data.Events.Image,
			Name:        data.Events.Name,
			Address:     data.Events.Address,
			Date:        data.Events.Date,
			Price:       data.Events.Price,
			Quote:       data.Events.Quote,
			Longitude:   data.Events.Longitude,
			Latitude:    data.Events.Latitude,
			Link:        data.Events.Link,
			Description: data.Events.Description,
			Status:      data.Events.Status,
		},
		Users: users.Core{
			ID:       int(data.User.ID),
			Image:    data.User.Email,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
	}
}
