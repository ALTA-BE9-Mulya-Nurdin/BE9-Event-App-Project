package response

import (
	_login "be9/event/features/login"
)

type User struct {
	Email string `json:"email"`
}

func FromCore(data _login.Core) User {
	return User{
		Email: data.Email,
	}
}
