package data

import (
	"be9/event/features/users"
	"fmt"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) InsertData(insert users.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) GetAllData() (data []users.Core, err error) {
	var getAllData []User
	tx := repo.db.Find(&getAllData)
	if tx.Error != nil {
		return []users.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}
