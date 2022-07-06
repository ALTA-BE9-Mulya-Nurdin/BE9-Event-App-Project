package comments

import "time"

type Core struct {
	ID          int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
	Event       Event
}

type User struct {
	ID       int
	Username string
}

type Event struct {
	ID int
}

type Business interface {
	GetDataAll() (data []Core, err error)
	InsertComment(insert Core) (row int, err error)
}

type Data interface {
	GetDataAll() (data []Core, err error)
	InsertComment(insert Core) (row int, err error)
}
