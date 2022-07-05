package data

import (
	_login "be9/event/features/login"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) _login.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}

func (repo *mysqlAuthRepository) Auth(data _login.Core) (dataAuth _login.Core, err error) {
	var authUser User
	tx := repo.db.First(&authUser, "email = ?", data.Email)
	if tx.Error != nil {
		return _login.Core{}, tx.Error
	}
	err = bcrypt.CompareHashAndPassword([]byte(authUser.toCore().Password), []byte(data.Password))
	if err != nil {
		return _login.Core{}, err
	}
	return authUser.toCore(), nil
}
