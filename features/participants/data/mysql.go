package data

import (
	"be9/event/features/participants"

	"gorm.io/gorm"
)

type mysqlParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(conn *gorm.DB) participants.Data {
	return &mysqlParticipantRepository{
		db: conn,
	}
}
