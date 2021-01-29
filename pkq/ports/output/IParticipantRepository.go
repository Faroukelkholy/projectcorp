package output

import "projectcorp/pkq/domain/model"

type IParticipantRepository interface {
	CreateParticipant(participant *model.Participant) error
}
