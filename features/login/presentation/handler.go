package presentation

import (
	_login "be9/event/features/login"
	"be9/event/features/login/presentation/request"
	"be9/event/features/login/presentation/response"
	"be9/event/helper"
	_middlewares "be9/event/middlewares"

	"github.com/go-playground/validator/v10"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authBusiness _login.Business
}

func NewAuthHandler(business _login.Business) *AuthHandler {
	return &AuthHandler{
		authBusiness: business,
	}
}

func (a *AuthHandler) Auth(c echo.Context) error {
	insertLogin := request.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errBind := c.Bind(&insertLogin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(insertLogin)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	authUser := request.ToCore(insertLogin)
	dataAuth, err := a.authBusiness.Auth(authUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error login"))
	}
	token, errToken := _middlewares.CreateToken(int(dataAuth.ID))
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error toker"))
	}
	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"message": "success",
			"data":    response.FromCore(dataAuth),
			"token":   token,
		},
	)

}
