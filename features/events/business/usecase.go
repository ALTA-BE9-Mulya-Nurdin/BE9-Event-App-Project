package business

import (
	"be9/event/features/events"
	"errors"
	"time"
)

type eventUsecase struct {
	eventData events.Data
}

func NewEventBusiness(evntData events.Data) events.Business {
	return &eventUsecase{
		eventData: evntData,
	}
}

func (uc *eventUsecase) InsertEvent(data events.Core) (int, error) {
	if data.Image == "" || data.Name == "" || data.Address == "" || data.Date == time.Now() || data.Price == 0 ||
		data.Quota == 0 || data.Longitude == "" || data.Latitude == "" || data.Link == "" || data.Description == "" || data.Status == "" {
		return -1, errors.New("failed to create event")
	}
	row, err := uc.eventData.InsertEvent(data)
	return row, err
}
