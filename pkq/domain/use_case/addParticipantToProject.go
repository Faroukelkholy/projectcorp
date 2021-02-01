package usecase

import (
	"errors"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/port/input/project"
	"projectcorp/pkq/port/output"
)

type AddParticipantToProject struct {
	proRepo output.IProjectRepository
	parRepo output.IParticipantRepository
}

func NewAddParticipantToProject(proRepo output.IProjectRepository, parRepo output.IParticipantRepository) projectPorts.IAddParticipantToProject {
	return &AddParticipantToProject{proRepo, parRepo}
}

func (ap *AddParticipantToProject) AddParticipant(participant *model.Participant) error {
	pro, errPro := ap.proRepo.GetProject(participant.ProjectID)
	if errPro != nil {
		return errors.New("no project exists for this id")
	}

	if pro.Department != participant.Department {
		return errors.New("project owner is not working at the same department as the participant")
	}

	return ap.parRepo.CreateParticipant(participant)
}
