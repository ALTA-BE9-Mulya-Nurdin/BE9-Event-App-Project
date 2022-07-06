package data

import (
	"be9/event/features/history"

	"gorm.io/gorm"
)

type mysqlHistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(conn *gorm.DB) history.Data {
	return &mysqlHistoryRepository{
		db: conn,
	}

}

func (repo *mysqlHistoryRepository) GetData(idToken int) (data []history.Core, err error) {
	var getHisEvents []Participants
	tx := repo.db.Where("user_id = ?", idToken).Preload("User").Preload("Events").Find(&getHisEvents)
	if tx.Error != nil {
		return []history.Core{}, tx.Error
	}
	return toCoreList(getHisEvents), nil
}
