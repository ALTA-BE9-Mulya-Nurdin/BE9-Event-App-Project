package business

import (
	_login "be9/event/features/login"
)

type authUsecase struct {
	authData _login.Data
}

func NewAuthBusiness(authData _login.Data) _login.Business {
	return &authUsecase{
		authData: authData,
	}
}

func (ac *authUsecase) Auth(data _login.Core) (dataAuth _login.Core, err error) {
	dataAuth, err = ac.authData.Auth(data)
	return dataAuth, err
}
