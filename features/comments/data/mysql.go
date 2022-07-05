package data

import (
	"be9/event/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		db: conn,
	}
}

func (repo *mysqlCommentRepository) GetDataAll() (data []comments.Core, err error) {
	var getAllComment []Comments
	tx := repo.db.Find(&getAllComment)
	if tx.Error != nil {
		return []comments.Core{}, tx.Error
	}
	return toCoreList(getAllComment), nil
}
