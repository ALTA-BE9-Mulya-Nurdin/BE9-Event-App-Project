package data

import (
	"be9/event/features/events"
	"fmt"
	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func (repo *mysqlEventRepository) InsertEventData(input events.Core) (int, error) {
	event := fromCore(input)

	res := *repo.db.Create(&event)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert input")
	}
	return int(res.RowsAffected), nil
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}

}
