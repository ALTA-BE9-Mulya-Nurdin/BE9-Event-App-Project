package factory

import (
	_authBusiness "be9/event/features/login/business"
	_authData "be9/event/features/login/data"
	_authPresentation "be9/event/features/login/presentation"

	eventBusiness "be9/event/features/events/business"
	eventData "be9/event/features/events/data"
	eventPresentation "be9/event/features/events/persentation"
	_userBusiness "be9/event/features/users/business"
	_userData "be9/event/features/users/data"
	_userPresentation "be9/event/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	// Login
	AuthPresenter *_authPresentation.AuthHandler
	// Users
	UserPresenter *_userPresentation.UserHandler
	// Event
	EventPresente *eventPresentation.EventHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// Users
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	eventData := eventData.NewEventRepository(dbConn)
	eventBusiness := eventBusiness.NewEventBusiness(eventData)
	eventPresentation := eventPresentation.NewEventHandler(eventBusiness)

	return Presenter{
		AuthPresenter: authPresentation,
		UserPresenter: userPresentation,
		EventPresente: eventPresentation,
	}
}
