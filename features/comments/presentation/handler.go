package presentation

import (
	"be9/event/features/comments"
	"be9/event/features/comments/presentation/request"
	"be9/event/features/comments/presentation/response"
	"be9/event/helper"
	_middlewares "be9/event/middlewares"
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

func (h *CommentHandler) InsertComment(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}

	var insertComment request.Comments
	errBind := c.Bind(&insertComment)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}

	newComment := request.ToCore(insertComment)
	newComment.User.ID = idToken
	row, err := h.commentBusiness.InsertComment(newComment)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert comment"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert comment"))
}
