package participants

import (
	"be9/event/features/events"
	"be9/event/features/users"
	"time"
)

type Core struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      users.Core
	Events    events.Core
}

type Business interface {
	GetDataAll() (data []Core, err error)
	GetData(idToken int, insert Core) (row int, err error)
	InsertData(insert Core) (row int, err error)
}

type Data interface {
	GetDataAll() (data []Core, err error)
	GetData(idToken int, insert Core) (row int, err error)
	InsertData(insert Core) (row int, err error)
}
