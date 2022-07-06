package business

import (
	"be9/event/features/participants"
)

type participantUseCase struct {
	participantData participants.Data
}

func NewParticipantBusiness(pdData participants.Data) participants.Business {
	return &participantUseCase{
		participantData: pdData,
	}
}

func (pd *participantUseCase) GetData(idToken int, insert participants.Core) (row int, err error) {
	row, err = pd.participantData.GetData(idToken, insert)
	return row, err
}

func (pd *participantUseCase) InsertData(insert participants.Core) (row int, err error) {
	row, err = pd.participantData.InsertData(insert)
	return row, err
}

func (pd *participantUseCase) GetDataAll() (data []participants.Core, err error) {
	data, err = pd.participantData.GetDataAll()
	return data, err
}
