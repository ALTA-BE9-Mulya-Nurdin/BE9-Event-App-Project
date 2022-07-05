package data

import (
	"be9/event/features/eventdetail"

	"gorm.io/gorm"
)

type mysqlEventDetailRepository struct {
	db *gorm.DB
}

func NewEventDetailRepository(conn *gorm.DB) eventdetail.Data {
	return &mysqlEventDetailRepository{
		db: conn,
	}
}
