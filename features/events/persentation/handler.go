package persentation

import (
	"be9/event/features/events"
	"be9/event/features/events/persentation/request"
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"net/http"
	"strings"
	"time"
)

var storageClient *storage.Client

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
	c.Bind(&newEvent)
	bucket := "event"

	var err error
	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("service-account-file.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Connect")
	}

	f, uploaded_file, err := c.Request().FormFile("photo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to Upload File")
	}

	defer f.Close()

	ext := strings.Split(uploaded_file.Filename, ".")
	extension := ext[len(ext)-1]
	check_extension := strings.ToLower(extension)
	if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
		return c.JSON(http.StatusBadRequest, "File Extension Not Allowed")
	}

	if uploaded_file.Size == 0 {
		return c.JSON(http.StatusBadRequest, "Illegal File")
	} else if uploaded_file.Size > 1050000 {
		return c.JSON(http.StatusBadRequest, "Size File Too Big")
	}
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	photo_name := strings.ReplaceAll(newEvent.Name, " ", "-")
	uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", photo_name, formatted, extension)
	newEvent.Image = uploaded_file.Filename
	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataEvent := request.ToCore(newEvent)
	row, err := handle.eventBisnis.InsertEvent(dataEvent)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "sukses insert data event")

}
