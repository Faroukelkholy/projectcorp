package mock

import "projectcorp/pkq/domain/model"

type ParticipantRepo struct {
	CreateParticipantMock func(participant *model.Participant) error
}

func (pr *ParticipantRepo) CreateParticipant(participant *model.Participant) error {
	return pr.CreateParticipantMock(participant)
}
