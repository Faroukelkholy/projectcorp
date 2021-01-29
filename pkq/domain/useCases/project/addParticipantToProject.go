package project

import (
	"errors"
	"log"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/ports/input/project"
	"projectcorp/pkq/ports/output"
)

type AddParticipantToProjectUseCase struct {
	projectRepo     output.IProjectRepository
	participantRepo output.IParticipantRepository
}

func NewAddParticipantToProjectUseCase(projectRepo output.IProjectRepository,participantRepo output.IParticipantRepository) projectPorts.IAddParticipantToProject {
	return &AddParticipantToProjectUseCase{projectRepo,participantRepo}
}

func (thisAP *AddParticipantToProjectUseCase) AddParticipant(participant *model.Participant) error {
	project, errorProjectRepo := thisAP.projectRepo.GetProject(participant.Project_id)
	if errorProjectRepo != nil {
		log.Println("error in AddParticipant UseCase :",errorProjectRepo)
		return errors.New("no project exists for this id")
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
