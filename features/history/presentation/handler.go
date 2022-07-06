package presentation

import (
	"be9/event/features/history"
	"be9/event/features/history/presentation/response"
	"be9/event/helper"
	_middlewares "be9/event/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HistoryHandler struct {
	historyBusiness history.Business
}

func NewHistoryHandler(business history.Business) *HistoryHandler {
	return &HistoryHandler{
		historyBusiness: business,
	}
}

func (p *HistoryHandler) GetData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, err := p.historyBusiness.GetData(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
