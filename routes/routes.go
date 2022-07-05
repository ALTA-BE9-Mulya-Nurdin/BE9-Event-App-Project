package routes

import (
	"be9/event/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	// Routes
	// Users
	e.POST("/users", presenter.UserPresenter.InsertData)
	e.GET("/users", presenter.UserPresenter.GetAllData)
	e.GET("/users/:id", presenter.UserPresenter.GetData)
	e.DELETE("/users/:id", presenter.UserPresenter.DeleteData)
	e.PUT("/users/:id", presenter.UserPresenter.UpdateData)

	return e
}
