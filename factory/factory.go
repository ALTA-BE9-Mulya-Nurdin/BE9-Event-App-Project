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

	_eventBusiness "be9/event/features/events/business"
	_eventData "be9/event/features/events/data"
	_eventPresentation "be9/event/features/events/presentation"

	_participantBusiness "be9/event/features/participants/business"
	_participantData "be9/event/features/participants/data"
	_participantPresentation "be9/event/features/participants/presentation"

	_historyBusiness "be9/event/features/history/business"
	_historyData "be9/event/features/history/data"
	_historyPresentation "be9/event/features/history/presentation"

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
	// Events
	EventPresenter *_eventPresentation.EventHandler
	// Participant
	ParticipantPresenter *_participantPresentation.ParticipantHandler
	// history
	HistoryPresenter *_historyPresentation.HistoryHandler
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

	eventData := _eventData.NewEventRepository(dbConn)
	eventBusiness := _eventBusiness.NewEventBusiness(eventData)
	eventPresentation := _eventPresentation.NewEventHandler(eventBusiness)

	participantData := _participantData.NewParticipantRepository(dbConn)
	participantBusiness := _participantBusiness.NewParticipantBusiness(participantData)
	participantPresentation := _participantPresentation.NewParticipantHandler(participantBusiness)

	historyData := _historyData.NewHistoryRepository(dbConn)
	historyBusiness := _historyBusiness.NewHistoryBusiness(historyData)
	historyPresentation := _historyPresentation.NewHistoryHandler(historyBusiness)

	return Presenter{
		AuthPresenter:        authPresentation,
		UserPresenter:        userPresentation,
		EventPresenter:       eventPresentation,
		CategorysPresenter:   categoryPresentation,
		CommentPresenter:     commentPresentation,
		ParticipantPresenter: participantPresentation,
		HistoryPresenter:     historyPresentation,
	}
}
