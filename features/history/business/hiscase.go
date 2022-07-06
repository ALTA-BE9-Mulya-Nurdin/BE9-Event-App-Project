package business

import (
	// "be9/event/features/events"
	"be9/event/features/history"
)

type historyUseCase struct {
	historyData history.Data
}

func NewHistoryBusiness(pdData history.Data) history.Business {
	return &historyUseCase{
		historyData: pdData,
	}
}

func (hiscase *historyUseCase) GetData(idToken int) (data []history.Core, err error) {
	data, err = hiscase.historyData.GetData(idToken)
	return data, err
}
