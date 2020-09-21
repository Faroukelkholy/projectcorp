package repositoryMock

import "projectcorp/domain/model"

type ParticipantRepoMock struct {
	CreateParticipantMock func(participant *model.Participant) error
}


func(thisPR *ParticipantRepoMock) CreateParticipant(participant *model.Participant) error{
	return thisPR.CreateParticipantMock(participant)
}
