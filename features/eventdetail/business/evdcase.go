package business

import "be9/event/features/eventdetail"

type eventDetailUsecase struct {
	eventDetailData eventdetail.Data
}

func NewEventDetailBusiness(evdData eventdetail.Data) eventdetail.Business {
	return &eventDetailUsecase{
		eventDetailData: evdData,
	}
}
