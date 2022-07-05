package data

import (
	"be9/event/features/categorys"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) categorys.Data {
	return &mysqlCategoryRepository{
		db: conn,
	}
}

func (repo *mysqlCategoryRepository) GetDataAll() (data []categorys.Core, err error) {
	var getAllCategory []Categorys
	tx := repo.db.Find(&getAllCategory)
	if tx.Error != nil {
		return []categorys.Core{}, tx.Error
	}
	return toCoreList(getAllCategory), nil
}
