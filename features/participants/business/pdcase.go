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
