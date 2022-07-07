package data

import (
	"be9/event/features/comments"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Description string `json:"description" form:"description"`
	UserID      uint   `json:"user_id" form:"user_id"`
	EventsID    uint   `json:"events_id" form:"events_id"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Events      Events `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Image    string `json:"image" form:"image"`
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Comments []Comments
}

type Events struct {
	gorm.Model
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	Date        string `json:"date" form:"date"`
	Price       int    `json:"price" form:"price"`
	Quote       int    `json:"quote" form:"quote"`
	Longitude   string `json:"longitude" form:"longitude"`
	Latitude    string `json:"latitude" form:"latitude"`
	Link        string `json:"link" form:"link"`
	Description string `json:"description" form:"description"`
	Status      string `json:"status" form:"status"`
	UserID      uint   `json:"user_id" form:"user_id"`
	CategorysID uint   `json:"categorys_id" form:"categorys_id"`
	Comments    []Comments
}

func (data *Comments) toCore() comments.Core {
	return comments.Core{
		ID:          int(data.ID),
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: comments.User{
			ID:       int(data.User.ID),
			Username: data.User.Username,
		},
		Events: comments.Events{
			ID:   int(data.Events.ID),
			Name: data.Events.Name,
		},
	}
}

func toCoreList(data []Comments) []comments.Core {
	result := []comments.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core comments.Core) Comments {
	return Comments{
		Description: core.Description,
		UserID:      uint(core.User.ID),
		EventsID:    uint(core.Events.ID),
	}
}
