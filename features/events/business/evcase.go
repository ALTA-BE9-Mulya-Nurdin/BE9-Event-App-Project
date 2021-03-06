package business

import "be9/event/features/events"

type eventUsecase struct {
	eventData events.Data
}

func NewEventBusiness(evData events.Data) events.Business {
	return &eventUsecase{
		eventData: evData,
	}
}

func (evcase *eventUsecase) InsertData(insert events.Core) (row int, err error) {
	row, err = evcase.eventData.InsertData(insert)
	return row, err
}

func (evcase *eventUsecase) GetAllData(limit int, offset int) (data []events.Core, err error) {
	data, err = evcase.eventData.GetAllData(limit, offset)
	return data, err
}

func (evcase *eventUsecase) GetData(id int) (data events.Core, err error) {
	data, err = evcase.eventData.GetData(id)
	return data, err
}

func (evcase *eventUsecase) DeleteData(id int) (row int, err error) {
	row, err = evcase.eventData.DeleteData(id)
	return row, err
}

func (evcase *eventUsecase) UpdatedData(id int, insert events.Core) (row int, err error) {
	row, err = evcase.eventData.UpdatedData(id, insert)
	return row, err
}

func (evcase *eventUsecase) GetToken(id int, idToken int) (data events.Core, err error) {
	data, err = evcase.eventData.GetToken(id, idToken)
	return data, err
}

func (evcase *eventUsecase) GetDataUserEvent(idToken int) (data []events.Core, err error) {
	data, err = evcase.eventData.GetDataUserEvent(idToken)
	return data, err
}

func (evcase *eventUsecase) GetDataHistoryEvent(idToken int) (data []events.Core, err error) {
	data, err = evcase.eventData.GetDataUserEvent(idToken)
	return data, err
}
