package project

import (
	"log"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	"projectcorp/ports/output"
)

type GetProjectsUseCase struct {
	projectRepo output.IProjectRepository
}


func NewGetProjectsUseCase(projectRepo output.IProjectRepository) projectPorts.IGetProjects {
	return &GetProjectsUseCase{projectRepo}
}

func (thisGP *GetProjectsUseCase) GetProjects() ([]*model.Project, error) {
	projects, err := thisGP.projectRepo.GetProjects()
	if err != nil {
		log.Println("error in GetPorjects UseCase :",err)
		return nil,err
	}
	return  projects, nil
}