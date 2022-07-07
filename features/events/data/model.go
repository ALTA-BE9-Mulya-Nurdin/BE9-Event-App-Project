package data

import (
	"be9/event/features/categorys"
	_categorys "be9/event/features/categorys/data"
	"be9/event/features/events"
	"be9/event/features/users"
	_users "be9/event/features/users/data"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	Image       string               `json:"image" form:"image"`
	Name        string               `json:"name" form:"name"`
	Address     string               `json:"address" form:"address"`
	Date        string               `json:"date" form:"date"`
	Price       int                  `json:"price" form:"price"`
	Quote       int                  `json:"quote" form:"quote"`
	Longitude   string               `json:"longitude" form:"longitude"`
	Latitude    string               `json:"latitude" form:"latitude"`
	Link        string               `json:"link" form:"link"`
	Description string               `json:"description" form:"description"`
	Status      string               `json:"status" form:"status"`
	UserID      uint                 `json:"user_id" form:"user_id"`
	CategorysID uint                 `json:"categorys_id" form:"categorys_id"`
	User        _users.User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categorys   _categorys.Categorys `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func toCoreList(data []Events) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Events) toCore() events.Core {
	return events.Core{
		ID:          int(data.ID),
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
		UpdatedAt:   data.UpdatedAt,
		User: users.Core{
			ID:       int(data.User.ID),
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Categorys: categorys.Core{
			ID:           int(data.Categorys.ID),
			CategoryName: data.Categorys.CategoryName,
		},
	}
}

func fromCore(core events.Core) Events {
	return Events{
		Image:       core.Image,
		Name:        core.Name,
		Address:     core.Address,
		Date:        core.Date,
		Price:       core.Price,
		Quote:       core.Quote,
		Longitude:   core.Longitude,
		Latitude:    core.Latitude,
		Link:        core.Link,
		Description: core.Description,
		Status:      core.Status,
		UserID:      uint(core.User.ID),
		CategorysID: uint(core.Categorys.ID),
	}
}
