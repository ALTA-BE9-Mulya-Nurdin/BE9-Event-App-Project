package response

import (
	_events "be9/event/features/events/presentation/response"
	"be9/event/features/participants"
	_users "be9/event/features/users/presentation/response"
	"time"
)

type Participants struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	User      _users.User
	Events    _events.Events
}

func FromCoreList(data []participants.Core) []Participants {
	result := []Participants{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data participants.Core) Participants {
	return Participants{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		User: _users.User{
			ID:       data.User.ID,
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Events: _events.Events{
			ID:          data.Events.ID,
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
