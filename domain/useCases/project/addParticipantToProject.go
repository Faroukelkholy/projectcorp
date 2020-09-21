package project

import (
	"errors"
	"log"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	"projectcorp/ports/output"
)

type AddParticipantToProjectUseCase struct {
	projectRepo output.IProjectRepository
	participantRepo output.IParticipantRepository
}

func NewAddParticipantToProjectUseCase(projectRepo output.IProjectRepository,participantRepo output.IParticipantRepository) projectPorts.IAddParticipantToProject {
	return &AddParticipantToProjectUseCase{projectRepo,participantRepo}
}

func (thisAP *AddParticipantToProjectUseCase) AddParticipant(participant *model.Participant) error {
	project, errorProjectRepo := thisAP.projectRepo.GetProject(participant.Project_id)
	if errorProjectRepo != nil {
		log.Println("error in AddParticipant UseCase :",errorProjectRepo)
		return errorProjectRepo
	}

	if project.Department != participant.Department {
		return errors.New("project owner is not working at the same department as the participant")
	}
	
	error :=thisAP.participantRepo.CreateParticipant(participant)

	if error != nil {
		log.Println("error in AddParticipant UseCase :",error)
		return error
	}

	return nil
}
