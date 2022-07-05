package events

import "time"

type Core struct {
	ID          int
	Image       string
	Name        string
	Address     string
	Date        time.Time
	Price       int
	Quota       int
	Longitude   string
	Latitude    string
	Link        string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	InsertEvent(data Core) (int, error)
}

type Data interface {
	InsertEvent(data Core) (int, error)
}
