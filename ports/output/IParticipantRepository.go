package output

import "projectcorp/domain/model"

type IParticipantRepository interface {
	CreateParticipant(participant *model.Participant) error
}
