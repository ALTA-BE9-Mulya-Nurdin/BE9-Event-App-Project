package routes

import (
	"be9/event/factory"
	_middlewares "be9/event/middlewares"

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

	// Login
	e.POST("/login", presenter.AuthPresenter.Auth)

	// Categories
	e.GET("/categorys", presenter.CategorysPresenter.GetDataAll)

	// Comments
	e.GET("/comments", presenter.CommentPresenter.GetDataAll)
	// Events
	e.POST("/events", presenter.EventPresenter.InsertData, _middlewares.JWTMiddleware())
	e.GET("/events", presenter.EventPresenter.GetAllData, _middlewares.JWTMiddleware())
	e.GET("/events/:id", presenter.EventPresenter.GetData, _middlewares.JWTMiddleware())
	e.DELETE("events/:id", presenter.EventPresenter.DeleteData, _middlewares.JWTMiddleware())
	e.PUT("/events/:id", presenter.EventPresenter.UpdateData, _middlewares.JWTMiddleware())

	return e
}
