package categorys

import "time"

type Core struct {
	ID           int
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Business interface {
	GetDataAll() (data []Core, err error)
}

type Data interface {
	GetDataAll() (data []Core, err error)
}
