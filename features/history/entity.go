package history

import (
	"be9/event/features/events"
	"be9/event/features/users"
	"time"
)

type Core struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     users.Core
	Events    events.Core
}

type Business interface {
	GetData(idToken int) (data []Core, err error)
}

type Data interface {
	GetData(idToken int) (data []Core, err error)
}
