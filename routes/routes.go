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
	e.GET("/users", presenter.UserPresenter.GetAllData, _middlewares.JWTMiddleware())
	e.GET("/users/profile", presenter.UserPresenter.GetData, _middlewares.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.DeleteData, _middlewares.JWTMiddleware())
	e.PUT("/users", presenter.UserPresenter.UpdateData, _middlewares.JWTMiddleware())

	// Login
	e.POST("/login", presenter.AuthPresenter.Auth)

	// Categories
	e.GET("/categorys", presenter.CategorysPresenter.GetDataAll)

	// Comments
	e.POST("/comments", presenter.CommentPresenter.InsertComment, _middlewares.JWTMiddleware())
	e.GET("/comments", presenter.CommentPresenter.GetDataAll)

	// Events
	e.POST("/events", presenter.EventPresenter.InsertData, _middlewares.JWTMiddleware())
	e.GET("/events", presenter.EventPresenter.GetAllData)
	e.GET("/events/:id", presenter.EventPresenter.GetData, _middlewares.JWTMiddleware())
	e.DELETE("events/:id", presenter.EventPresenter.DeleteData, _middlewares.JWTMiddleware())
	e.PUT("/events/:id", presenter.EventPresenter.UpdateData, _middlewares.JWTMiddleware())
	e.GET("/user/events", presenter.EventPresenter.GetDataUserEvent, _middlewares.JWTMiddleware())

	e.GET("/history_events", presenter.HistoryPresenter.GetData, _middlewares.JWTMiddleware())

	// Participants
	e.POST("/participants", presenter.ParticipantPresenter.InsertData, _middlewares.JWTMiddleware())
	e.GET("/participants", presenter.ParticipantPresenter.GetDataAll)

	return e
}
