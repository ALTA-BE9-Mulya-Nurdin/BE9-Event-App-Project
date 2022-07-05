package factory

import (
	_userBusiness "be9/event/features/users/business"
	_userData "be9/event/features/users/data"
	_userPresentation "be9/event/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	// Users
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// Users
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		// Users
		UserPresenter: userPresentation,
	}
}
