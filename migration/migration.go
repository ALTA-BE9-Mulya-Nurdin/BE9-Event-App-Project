package migration

import (
	_mCategorys "be9/event/features/categorys/data"
	_mEvents "be9/event/features/events/data"
	_mUsers "be9/event/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mCategorys.Categorys{})
	db.AutoMigrate(&_mEvents.Events{})
}
