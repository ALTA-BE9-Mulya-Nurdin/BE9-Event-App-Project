package request

import "be9/event/features/users"

type User struct {
	Image    string `json:"image" validate:"required" form:"image"`
	Username string `json:"username" validate:"required" form:"username"`
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
	Phone    string `json:"phone" validate:"required,numeric" form:"phone"`
	Address  string `json:"address" validate:"required" form:"address"`
}

func ToCore(req User) users.Core {
	return users.Core{
		Image:    req.Image,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
	}
}
