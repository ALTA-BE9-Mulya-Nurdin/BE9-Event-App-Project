package data

import (
	_login "be9/event/features/login"
)

type User struct {
	ID       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (data *User) toCore() _login.Core {
	return _login.Core{
		ID:       int(data.ID),
		Email:    data.Email,
		Password: data.Password,
	}
}

func fromCore(core _login.Core) User {
	return User{
		Email:    core.Email,
		Password: core.Password,
	}
}
