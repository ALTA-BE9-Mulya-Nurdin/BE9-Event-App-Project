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

// func (repo *mysqlEventDetailRepository) GetData(idToken int, idEvent int) (row int, err error) {
// 	var getData EventDetail
// 	tx := repo.db.Where("user_id = ? AND events_id = ?", idToken, idEvent).First(&getData)
// }
