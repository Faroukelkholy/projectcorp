package output

import "projectcorp/domain/model"

type IProjectRepository interface {
	GetProjects() ([]*model.Project,error)
	GetProject(id string) (*model.Project,error)
	CreateProject(project *model.Project) error
	UpdateProject(project *model.Project) error
}
