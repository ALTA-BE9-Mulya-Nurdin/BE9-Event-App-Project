package response

import (
	_events "be9/event/features/events/presentation/response"
	"be9/event/features/history"
	_users "be9/event/features/users/presentation/response"
	"time"
)

type Historys struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	User      _users.User
	Events    _events.Events
}

func FromCoreList(data []history.Core) []Historys {
	result := []Historys{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data history.Core) Historys {
	return Historys{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		User: _users.User{
			ID:       data.Users.ID,
			Image:    data.Users.Image,
			Username: data.Users.Username,
			Email:    data.Users.Email,
			Phone:    data.Users.Phone,
			Address:  data.Users.Address,
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
