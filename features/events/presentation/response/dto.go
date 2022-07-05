package response

import (
	_categorys "be9/event/features/categorys/presentation/response"
	"be9/event/features/events"
	_users "be9/event/features/users/presentation/response"
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
	User        _users.User
	Categorys   _categorys.Categorys
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
		User: _users.User{
			ID:       data.User.ID,
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Categorys: _categorys.Categorys{
			ID:           data.Categorys.ID,
			CategoryName: data.Categorys.CategoryName,
		},
	}
}
