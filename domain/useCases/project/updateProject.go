package project


import (
	"errors"
	"log"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	"projectcorp/ports/output"
)

type UpdateProjectUseCase struct {
	projectRepo output.IProjectRepository
	restClient output.IClientRest
}


func NewUpdateProjectUseCase(projectRepo output.IProjectRepository,restClient output.IClientRest) projectPorts.IUpdateProject {
	return &UpdateProjectUseCase{projectRepo,restClient}
}

func (thisUP *UpdateProjectUseCase) UpdateProject(project *model.Project) error {
	employee,errorClientRest := thisUP.restClient.GetEmployee(config.GETEMPLOYEES_URL,project.Owner)

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

