package usecase

import (
	"errors"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/port/input/project"
	"projectcorp/pkq/port/output"
)

type UpdateProject struct {
	proRepo output.IProjectRepository
	hc      output.IHTTPClient
}

func NewUpdateProject(proRepo output.IProjectRepository, hc output.IHTTPClient) projectPorts.IUpdateProject {
	return &UpdateProject{proRepo, hc}
}

func (up *UpdateProject) UpdateProject(project *model.Project) error {
	e, errorHC := up.hc.GetEmployee(project.Owner)
	if errorHC != nil {
		return errorHC
	}

	if e.Role != "manager" {
		return errors.New("project owner is not a manager")
	}

	return up.proRepo.UpdateProject(project)
}
