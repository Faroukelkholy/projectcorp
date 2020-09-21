package project

import "projectcorp/domain/model"
type IAddParticipantToProject interface {
	AddParticipant(participant *model.Participant) error
}