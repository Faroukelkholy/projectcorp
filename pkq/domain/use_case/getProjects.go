package usecase

import (
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/port/input/project"
	"projectcorp/pkq/port/output"
)

type GetProjects struct {
	proRepo output.IProjectRepository
}

func NewGetProjects(proRepo output.IProjectRepository) projectPorts.IGetProjects {
	return &GetProjects{proRepo}
}

func (gp *GetProjects) GetProjects() ([]*model.Project, error) {
	return gp.proRepo.GetProjects()
}
