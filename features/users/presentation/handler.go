package presentation

import (
	"be9/event/features/users"
	"be9/event/features/users/presentation/request"
	"be9/event/features/users/presentation/response"
	"be9/event/helper"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) InsertData(c echo.Context) error {
	var insertData request.User
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(insertData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newUser := request.ToCore(insertData)
	newUser.Image = "image.jpg"
	row, err := h.userBusiness.InsertData(newUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (h *UserHandler) GetAllData(c echo.Context) error {
	data, err := h.userBusiness.GetAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (h *UserHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idUser, _ := strconv.Atoi(id)
	data, err := h.userBusiness.GetData(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (h *UserHandler) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idUser, _ := strconv.Atoi(id)
	row, err := h.userBusiness.DeleteData(idUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}

func (h *UserHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idUser, _ := strconv.Atoi(id)
	data, _ := h.userBusiness.GetData(idUser)
	updatedData := request.User{
		Image:    c.FormValue("image"),
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Phone:    c.FormValue("phone"),
		Address:  c.FormValue("address"),
	}
	if updatedData.Image == "" {
		updatedData.Image = data.Image
	}
	if updatedData.Username == "" {
		updatedData.Username = data.Username
	}
	if updatedData.Email == "" {
		updatedData.Email = data.Email
	}
	if updatedData.Password == "" {
		updatedData.Password = data.Password
	} else if updatedData.Password != "" {
		hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if errHash != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to generate password"))
		}
		updatedData.Password = string(hashedPassword)
	}
	if updatedData.Phone == "" {
		updatedData.Phone = data.Phone
	}
	if updatedData.Address == "" {
		updatedData.Address = data.Address
	}
	v := validator.New()
	errValidator := v.Struct(updatedData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newUser := request.ToCore(updatedData)
	row, err := h.userBusiness.UpdateData(idUser, newUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}
