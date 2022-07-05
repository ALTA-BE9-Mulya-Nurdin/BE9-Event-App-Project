package presentation

import (
	"be9/event/features/events"
	"be9/event/features/events/presentation/request"
	"be9/event/features/events/presentation/response"
	"be9/event/helper"
	"net/http"
	"strconv"

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
	var insertData request.Events
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
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
	data, err := v.eventBusiness.GetAllData()
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
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
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
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	data, _ := v.eventBusiness.GetData(idEvent)
	priceInt, _ := strconv.Atoi(c.FormValue("price"))
	quoteInt, _ := strconv.Atoi(c.FormValue("quote"))
	updatedData := request.Events{
		Image:       c.FormValue("image"),
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
	if updatedData.Image == "" {
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
