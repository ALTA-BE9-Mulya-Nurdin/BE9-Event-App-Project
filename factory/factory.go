package factory

import (
	_commentBusiness "be9/event/features/comments/business"
	_commentData "be9/event/features/comments/data"
	_commentPresentation "be9/event/features/comments/presentation"

	_categoryBusiness "be9/event/features/categorys/business"
	_categoryData "be9/event/features/categorys/data"
	_categoryPresentation "be9/event/features/categorys/presentation"

	_authBusiness "be9/event/features/login/business"
	_authData "be9/event/features/login/data"
	_authPresentation "be9/event/features/login/presentation"

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
	// Categories
	CategorysPresenter *_categoryPresentation.CategoryHandler
	// Comments
	CommentPresenter *_commentPresentation.CommentHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// Users
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	categoryData := _categoryData.NewCategoryRepository(dbConn)
	categoryBusiness := _categoryBusiness.NewCategoryBusiness(categoryData)
	categoryPresentation := _categoryPresentation.NewCategoryHandler(categoryBusiness)

	commentData := _commentData.NewCommentRepository(dbConn)
	commentBusiness := _commentBusiness.NewCommentBusiness(commentData)
	commentPresentation := _commentPresentation.NewCommentHandler(commentBusiness)

	return Presenter{
		AuthPresenter:      authPresentation,
		UserPresenter:      userPresentation,
		CategorysPresenter: categoryPresentation,
		CommentPresenter:   commentPresentation,
	}
}
