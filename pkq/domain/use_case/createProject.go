package usecase

import (
	"errors"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/port/input/project"
	"projectcorp/pkq/port/output"
)

type CreateProject struct {
	proRepo output.IProjectRepository
	hc      output.IHTTPClient
}

func NewCreateProject(proRepo output.IProjectRepository, hc output.IHTTPClient) projectPorts.ICreateProject {
	return &CreateProject{proRepo, hc}
}

func (cp *CreateProject) CreateProject(project *model.Project) error {
	e, errorHC := cp.hc.GetEmployee(project.Owner)
	if errorHC != nil {
		return errorHC
	}

	if e.Role != "manager" {
		return errors.New("project owner is not a manager")
	}

	return cp.proRepo.CreateProject(project)
}
