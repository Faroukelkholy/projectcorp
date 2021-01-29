package project


import (
	"errors"
	"log"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/ports/input/project"
	"projectcorp/pkq/ports/output"
)

type UpdateProjectUseCase struct {
	projectRepo output.IProjectRepository
	restClient  output.IClientRest
}


func NewUpdateProjectUseCase(projectRepo output.IProjectRepository,restClient output.IClientRest) projectPorts.IUpdateProject {
	return &UpdateProjectUseCase{projectRepo,restClient}
}

func (thisUP *UpdateProjectUseCase) UpdateProject(project *model.Project) error {
	employee,errorClientRest := thisUP.restClient.GetEmployee(project.Owner)

	if errorClientRest != nil {
		log.Println("error in UpdateProject UseCase :",errorClientRest)
		return errorClientRest
	}
	if employee.Role != "manager" {
		return errors.New("project owner is not a manager")
	}

	err := thisUP.projectRepo.UpdateProject(project)
	if err != nil {
		log.Println("error in UpdateProject UseCase :",err)
		return err
	}
	return  nil
}

