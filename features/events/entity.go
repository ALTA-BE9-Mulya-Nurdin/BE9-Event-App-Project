package events

import (
	"be9/event/features/categorys"
	"be9/event/features/users"
	"time"
)

type Core struct {
	ID          int
	Image       string
	Name        string
	Address     string
	Date        string
	Price       int
	Quote       int
	Longitude   string
	Latitude    string
	Link        string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        users.Core
	Categorys   categorys.Core
}

type Business interface {
	InsertData(insert Core) (row int, err error)
	GetAllData() (data []Core, err error)
	GetData(id int) (data Core, err error)
	DeleteData(id int) (row int, err error)
	UpdatedData(id int, insert Core) (row int, err error)
	GetToken(id int, idToken int) (data Core, err error)
	GetDataUserEvent(idToken int) (data []Core, err error)
	GetDataHistoryEvent(idToken int) (data []Core, err error)
}

type Data interface {
	InsertData(insert Core) (row int, err error)
	GetAllData() (data []Core, err error)
	GetData(id int) (data Core, err error)
	DeleteData(id int) (row int, err error)
	UpdatedData(id int, insert Core) (row int, err error)
	GetToken(id int, idToken int) (data Core, err error)
	GetDataUserEvent(idToken int) (data []Core, err error)
	GetDataHistoryEvent(idToken int) (data []Core, err error)
}
