package comments

import "time"

type Core struct {
	ID          int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// User User
	// Event Event
}

// type User struct {
// 	ID   int
// }

// type Event struct {
// 	ID   int
// }

type Business interface {
	GetDataAll() (data []Core, err error)
}

type Data interface {
	GetDataAll() (data []Core, err error)
}
