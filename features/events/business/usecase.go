package business

import (
	"be9/event/features/events"
)

type eventUsecase struct {
	eventData events.Data
}

func NewEventBusiness(evntData events.Data) events.Business {
	return &eventUsecase{
		eventData: evntData,
	}
}

func (uc *eventUsecase) InsertEvent(data events.Core) (row int, err error) {
	row, err = uc.eventData.InsertEvent(data)
	return row, err
}

func (uc *eventUsecase) GetAllEvent() (data []events.Core, err error) {
	data, err = uc.eventData.GetAllEvent()
	return data, err
}

func (uc *eventUsecase) GetEvent(id int) (data events.Core, err error) {
	data, err = uc.eventData.GetEvent(id)
	return data, err
}

func (uc *eventUsecase) DeleteEvent(id int) (row int, err error) {
	row, err = uc.eventData.DeleteEvent(id)
	return row, err
}

func (uc *eventUsecase) UpdateEvent(id int, insert events.Core) (row int, err error) {
	row, err = uc.eventData.UpdateEvent(id, insert)
	return row, err
}
