package request

import (
	"be9/event/features/events"
)

type Event struct {
	//IDUsers int `json:"id_users" validate:"required" from:"id_users" `
	//IDCategory  int       `json:"id_category" validate:"required" from:"id_category"`
	Image       string `json:"image" validate:"required" from:"image"`
	Name        string `json:"name" validate:"required" from:"name"`
	Address     string `json:"address" validate:"required" from:"address"`
	Date        string `json:"date" validate:"required" from:"date"`
	Price       int    `json:"price" validate:"required" from:"price"`
	Quota       int    `json:"quota" validate:"required" from:"quota"`
	Longitude   string `json:"longitude" validate:"required" from:"longitude"`
	Latitude    string `json:"latitude" validate:"required" from:"latitude"`
	Link        string `json:"link" validate:"required" from:"link"`
	Description string `json:"description" validate:"required" from:"description"`
	Status      string `json:"status" validate:"required" from:"status"`
}

func ToCore(req Event) events.Core {
	return events.Core{
		//IDUsers: req.IDUsers,
		//IDCategory:  req.IDCategory,
		Image:       req.Image,
		Name:        req.Name,
		Address:     req.Address,
		Date:        req.Date,
		Price:       req.Price,
		Quota:       req.Quota,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Link:        req.Link,
		Description: req.Description,
		Status:      req.Status,
	}
}

//type User struct {
//	Image    string `json:"image" validate:"required" form:"image"`
//	Username string `json:"username" validate:"required" form:"username"`
//	Email    string `json:"email" validate:"required,email" form:"email"`
//	Password string `json:"password" validate:"required" form:"password"`
//	Phone    string `json:"phone" validate:"required,numeric" form:"phone"`
//	Address  string `json:"address" validate:"required" form:"address"`
//}
//
//func ToCore(req User) users.Core {
//	return users.Core{
//		Image:    req.Image,
//		Username: req.Username,
//		Email:    req.Email,
//		Password: req.Password,
//		Phone:    req.Phone,
//		Address:  req.Address,
//	}
//}
