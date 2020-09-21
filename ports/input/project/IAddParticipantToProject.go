package project

import "projectcorp/domain/model"
type IAddParticipantToProject interface {
	CreateProject(id string,participant *model.Participant) error
}