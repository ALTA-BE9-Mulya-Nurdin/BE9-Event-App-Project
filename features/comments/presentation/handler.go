package presentation

import (
	"be9/event/features/comments"
	"be9/event/features/comments/presentation/response"
	"be9/event/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(business comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: business,
	}
}

func (h *CommentHandler) GetDataAll(c echo.Context) error {
	data, err := h.commentBusiness.GetDataAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

// func (h *CommentHandler) InsertComment(c echo.Context) error {

// }
