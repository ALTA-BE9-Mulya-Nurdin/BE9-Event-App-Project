package presentation

import (
	"be9/event/features/categorys"
	"be9/event/features/categorys/presentation/response"
	"be9/event/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryBusiness categorys.Business
}

func NewCategoryHandler(business categorys.Business) *CategoryHandler {
	return &CategoryHandler{
		categoryBusiness: business,
	}
}

func (h *CategoryHandler) GetDataAll(c echo.Context) error {
	data, err := h.categoryBusiness.GetDataAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
