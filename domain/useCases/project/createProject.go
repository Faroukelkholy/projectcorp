package project

import (
	"errors"
	"log"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	"projectcorp/ports/output"
)

type CreateProjectUseCase struct {
	projectRepo output.IProjectRepository
	restClient output.IClientRest
}


func NewCreateProjectUseCase(projectRepo output.IProjectRepository,restClient output.IClientRest) projectPorts.ICreateProject {
	return &CreateProjectUseCase{projectRepo,restClient}
}

func (thisCP *CreateProjectUseCase) CreateProject(project *model.Project) error {
	employee,errorClientRest := thisCP.restClient.GetEmployee(project.Owner)

	if errorClientRest != nil {
		log.Println("error in CreateProject UseCase :",errorClientRest)
		return errorClientRest
	}
	if employee.Role != "manager" {
		return errors.New("project owner is not a manager")
	}

	err := thisCP.projectRepo.CreateProject(project)
	if err != nil {
		log.Println("error in CreateProject UseCase :",err)
		return err
	}
	return  nil
}
