package presentation

import (
	"be9/event/features/events"
	"be9/event/features/events/presentation/request"
	"be9/event/features/events/presentation/response"
	"be9/event/helper"
	_middlewares "be9/event/middlewares"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventBusiness events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		eventBusiness: business,
	}
}

func (v *EventHandler) InsertData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	link, report, err := helper.AddImageEvent(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var insertData request.Events
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	insertData.Image = link
	val := validator.New()
	errValidator := val.Struct(insertData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newEvent := request.ToCore(insertData)
	if newEvent.Longitude == "" {
		newEvent.Longitude = "NULL"
	}
	if newEvent.Latitude == "" {
		newEvent.Latitude = "NULL"
	}
	if newEvent.Link == "" {
		newEvent.Link = "NULL"
	}
	newEvent.User.ID = idToken
	row, err := v.eventBusiness.InsertData(newEvent)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (v *EventHandler) GetAllData(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	data, err := v.eventBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (v *EventHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	data, err := v.eventBusiness.GetData(idEvent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (v *EventHandler) DeleteData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	data, _ := v.eventBusiness.GetToken(idEvent, idToken)
	if data.User.ID != idToken {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	row, err := v.eventBusiness.DeleteData(idEvent)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}

func (v *EventHandler) UpdateData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	dataToken, _ := v.eventBusiness.GetToken(idEvent, idToken)
	if dataToken.User.ID != idToken {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	link, report, err := helper.AddImageEvent(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	data, _ := v.eventBusiness.GetData(idEvent)
	priceInt, _ := strconv.Atoi(c.FormValue("price"))
	quoteInt, _ := strconv.Atoi(c.FormValue("quote"))
	categorysIDInt, _ := strconv.Atoi(c.FormValue("categorys_id"))
	updatedData := request.Events{
		CategorysID: uint(categorysIDInt),
		Image:       link,
		Name:        c.FormValue("name"),
		Address:     c.FormValue("address"),
		Date:        c.FormValue("date"),
		Price:       priceInt,
		Quote:       quoteInt,
		Longitude:   c.FormValue("longitude"),
		Latitude:    c.FormValue("latitude"),
		Link:        c.FormValue("link"),
		Description: c.FormValue("description"),
		Status:      c.FormValue("status"),
	}
	if updatedData.CategorysID == 0 {
		updatedData.CategorysID = uint(data.Categorys.ID)
	}
	if updatedData.Image == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updatedData.Image = data.Image
	}
	if updatedData.Name == "" {
		updatedData.Name = data.Name
	}
	if updatedData.Address == "" {
		updatedData.Address = data.Address
	}
	if updatedData.Date == "" {
		updatedData.Date = data.Date
	}
	if updatedData.Price == 0 {
		updatedData.Price = data.Price
	}
	if updatedData.Quote == 0 {
		updatedData.Quote = data.Quote
	}
	if updatedData.Longitude == "" {
		updatedData.Longitude = data.Longitude
	}
	if updatedData.Latitude == "" {
		updatedData.Latitude = data.Latitude
	}
	if updatedData.Link == "" {
		updatedData.Link = data.Link
	}
	if updatedData.Description == "" {
		updatedData.Description = data.Description
	}
	if updatedData.Status == "" {
		updatedData.Status = data.Status
	}
	val := validator.New()
	errValidator := val.Struct(updatedData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newEvent := request.ToCore(updatedData)
	row, err := v.eventBusiness.UpdatedData(idEvent, newEvent)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}

func (v *EventHandler) GetDataUserEvent(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, err := v.eventBusiness.GetDataUserEvent(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
