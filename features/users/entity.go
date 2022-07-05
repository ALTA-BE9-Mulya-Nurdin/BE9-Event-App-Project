package users

import "time"

type Core struct {
	ID        int
	Image     string
	Username  string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
}

type Data interface {
}