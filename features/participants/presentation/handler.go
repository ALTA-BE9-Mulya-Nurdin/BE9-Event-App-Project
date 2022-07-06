package presentation

import (
	"be9/event/features/participants"
	"be9/event/features/participants/presentation/request"
	"be9/event/features/participants/presentation/response"
	"be9/event/helper"
	_middlewares "be9/event/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ParticipantHandler struct {
	participantBusiness participants.Business
}

func NewParticipantHandler(business participants.Business) *ParticipantHandler {
	return &ParticipantHandler{
		participantBusiness: business,
	}
}

func (p *ParticipantHandler) InsertData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	var insertData request.Participants
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	newData := request.ToCore(insertData)
	newData.User.ID = idToken
	row, _ := p.participantBusiness.GetData(idToken, newData)
	if row == 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("already joined the event"))
	}
	rowInsert, err := p.participantBusiness.InsertData(newData)
	if rowInsert != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to join event"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to join event"))
}

func (p *ParticipantHandler) GetDataAll(c echo.Context) error {
	data, err := p.participantBusiness.GetDataAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
