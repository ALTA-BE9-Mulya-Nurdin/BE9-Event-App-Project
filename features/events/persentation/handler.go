package persentation

import (
	"be9/event/features/events"
	"be9/event/features/events/persentation/request"
	"be9/event/features/events/persentation/response"
	"be9/event/helper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//var storageClient *storage.Client

type EventHandler struct {
	eventBisnis events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		eventBisnis: business,
	}
}

func (handle *EventHandler) CreateEvent(c echo.Context) error {
	//idUser := echo.MiddlewareFunc(c)
	//if idUser == 0 {
	//	return errors.New("failed")
	//}

	// buat event
	var newEvent request.Event
	errBind := c.Bind(&newEvent)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	//bucket := "event"

	//v := validator.New()
	//errValidator := v.Struct(newEvent)
	//if errValidator != nil {
	//	return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	//}

	//var err error
	//ctx := appengine.NewContext(c.Request())
	//
	//storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("service-account-file.json"))
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, "Can't Connect")
	//}
	//
	//f, uploaded_file, err := c.Request().FormFile("photo")
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, "Failed to Upload File")
	//}
	//
	//defer f.Close()
	//
	//ext := strings.Split(uploaded_file.Filename, ".")
	//extension := ext[len(ext)-1]
	//check_extension := strings.ToLower(extension)
	//if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
	//	return c.JSON(http.StatusBadRequest, "File Extension Not Allowed")
	//}
	//
	//if uploaded_file.Size == 0 {
	//	return c.JSON(http.StatusBadRequest, "Illegal File")
	//} else if uploaded_file.Size > 1050000 {
	//	return c.JSON(http.StatusBadRequest, "Size File Too Big")
	//}
	//t := time.Now()
	//formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
	//	t.Year(), t.Month(), t.Day(),
	//	t.Hour(), t.Minute(), t.Second())
	//photo_name := strings.ReplaceAll(newEvent.Name, " ", "-")
	//uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", photo_name, formatted, extension)
	//newEvent.Image = uploaded_file.Filename
	//sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)
	//
	//if _, err := io.Copy(sw, f); err != nil {
	//	return c.JSON(http.StatusInternalServerError, err.Error())
	//}
	//
	//if err := sw.Close(); err != nil {
	//	return c.JSON(http.StatusInternalServerError, err.Error())
	//}

	dataEvent := request.ToCore(newEvent)
	//dataEvent.IDUsers = 1
	//dataEvent.IDCategory = 1
	row, err := handle.eventBisnis.InsertEvent(dataEvent)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success")

}

func (handle *EventHandler) GetAllEvent(c echo.Context) error {
	data, err := handle.eventBisnis.GetAllEvent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (handle *EventHandler) GetEvent(c echo.Context) error {
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	data, err := handle.eventBisnis.GetEvent(idEvent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (handle *EventHandler) DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	row, err := handle.eventBisnis.DeleteEvent(idEvent)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data event"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data event"))
}

//func (handle *EventHandler) UpdateData(c echo.Context) error {
//	id := c.Param("id")
//	idEvent, _ := strconv.Atoi(id)
//	data, _ := handle.eventBisnis.GetEvent(idEvent)
//	updatedData := request.Event{
//		Image:       c.FormValue("image"),
//		Name:        c.FormValue("name"),
//		Address:     c.FormValue("address"),
//		Date:        c.FormValue("date"),
//		Price:       c.FormValue("price"),
//		Quota:       c.FormValue("quota"),
//		Longitude:   c.FormValue("longitude"),
//		Latitude:    c.FormValue("Latitude"),
//		Link:        c.FormValue("link"),
//		Description: c.FormValue("description"),
//		Status:      c.FormValue("status"),
//	}
//	if updatedData.Image == "" {
//		updatedData.Image = data.Image
//	}
//	if updatedData.Name == "" {
//		updatedData.Name = data.Name
//	}
//	if updatedData.Address == "" {
//		updatedData.Address = data.Address
//	}
//	if updatedData.Date == "" {
//		updatedData.Date = data.Date
//	}
//	if updatedData.Price == 0 {
//		updatedData.Price = data.Price
//	}
//	if updatedData.Quota == 0 {
//		updatedData.Quota = data.Quota
//	}
//	if updatedData.Longitude == "" {
//		updatedData.Longitude = data.Longitude
//	}
//	if updatedData.Latitude == "" {
//		updatedData.Latitude = data.Latitude
//	}
//	if updatedData.Link == "" {
//		updatedData.Link = data.Link
//	}
//	if updatedData.Description == "" {
//		updatedData.Description = data.Description
//	}
//	if updatedData.Status == "" {
//		updatedData.Status = data.Status
//	}
//	v := validator.New()
//	errValidator := v.Struct(updatedData)
//	if errValidator != nil {
//		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
//	}
//	newEvent := request.ToCore(updatedData)
//	row, err := handle.eventBisnis.UpdateEvent(idEvent, newEvent)
//	if row != 1 {
//		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data event"))
//	}
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
//	}
//	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data event"))
//}
