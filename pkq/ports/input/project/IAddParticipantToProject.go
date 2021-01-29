package project

import "projectcorp/pkq/domain/model"
type IAddParticipantToProject interface {
	AddParticipant(participant *model.Participant) error
}