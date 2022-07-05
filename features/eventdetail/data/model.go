package data

import (
	"be9/event/features/eventdetail"
	"be9/event/features/events"
	_events "be9/event/features/events/data"
	"be9/event/features/users"
	_users "be9/event/features/users/data"

	"gorm.io/gorm"
)

type EventDetail struct {
	gorm.Model
	UserID   uint `json:"user_id" form:"user_id"`
	EventsID uint `json:"events_id" form:"events_id"`
	User     _users.User
	Events   _events.Events
}

func toCoreList(data []EventDetail) []eventdetail.Core {
	result := []eventdetail.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *EventDetail) toCore() eventdetail.Core {
	return eventdetail.Core{
		ID:        int(data.ID),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: users.Core{
			ID:       int(data.User.ID),
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
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
	}
}

func fromCore(core eventdetail.Core) EventDetail {
	return EventDetail{
		UserID:   uint(core.User.ID),
		EventsID: uint(core.Events.ID),
	}
}
