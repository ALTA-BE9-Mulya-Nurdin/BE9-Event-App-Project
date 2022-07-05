package request

import (
	_login "be9/event/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

func ToCore(req User) _login.Core {
	return _login.Core{
		Email:    req.Email,
		Password: req.Password,
	}
}
