package events

import "time"

type Core struct {
	ID          int
	IDUsers     int
	IDCategory  int
	Image       string
	Name        string
	Address     string
	Date        string
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
	GetAllEvent() (data []Core, err error)
	GetEvent(id int) (data Core, err error)
	DeleteEvent(id int) (row int, err error)
	UpdateEvent(id int, insert Core) (row int, err error)
}

type Data interface {
	InsertEvent(data Core) (int, error)
	GetAllEvent() (data []Core, err error)
	GetEvent(id int) (data Core, err error)
	DeleteEvent(id int) (row int, err error)
	UpdateEvent(id int, insert Core) (row int, err error)
}
