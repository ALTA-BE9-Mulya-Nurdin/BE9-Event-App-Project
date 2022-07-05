package factory

import (
	_authBusiness "be9/event/features/login/business"
	_authData "be9/event/features/login/data"
	_authPresentation "be9/event/features/login/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	AuthPresenter *_authPresentation.AuthHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	return Presenter{
		AuthPresenter: authPresentation,
	}
}
