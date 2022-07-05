package migration

import (
	event "be9/event/features/events/data"
	_mUsers "be9/event/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&event.Event{})
}
