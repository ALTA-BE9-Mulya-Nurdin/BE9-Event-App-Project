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

func (repo *mysqlParticipantRepository) GetData(idToken int, insert participants.Core) (row int, err error) {
	var getData Participants
	tx := repo.db.Where("user_id = ? AND events_id = ?", idToken, insert.Events.ID).First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlParticipantRepository) InsertData(insert participants.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlParticipantRepository) GetDataAll() (data []participants.Core, err error) {
	var getDataAll []Participants
	tx := repo.db.Preload("User").Preload("Events").Order("created_at").Limit(5).Find(&getDataAll)
	if tx.Error != nil {
		return []participants.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}
