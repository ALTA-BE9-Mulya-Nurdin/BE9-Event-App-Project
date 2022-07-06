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
}

type Data interface {
}
