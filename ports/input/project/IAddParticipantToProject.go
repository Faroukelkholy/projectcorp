package project

import "projectcorp/domain/model"
type IAddParticipantToProject interface {
	AddParticipant(projectId string,participant *model.Participant) error
}