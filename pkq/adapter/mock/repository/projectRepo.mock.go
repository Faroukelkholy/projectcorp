package mock

import "projectcorp/pkq/domain/model"

type ProjectRepo struct {
	GetProjectsMock   func() ([]*model.Project, error)
	GetProjectMock    func(id string) (*model.Project, error)
	CreateProjectMock func(project *model.Project) error
	UpdateProjectMock func(project *model.Project) error
}

func (pr *ProjectRepo) GetProjects() ([]*model.Project, error) {
	return pr.GetProjectsMock()
}
func (pr *ProjectRepo) GetProject(id string) (*model.Project, error) {
	return pr.GetProjectMock(id)
}

func (pr *ProjectRepo) CreateProject(project *model.Project) error {
	return pr.CreateProjectMock(project)
}

func (pr *ProjectRepo) UpdateProject(project *model.Project) error {
	return pr.UpdateProjectMock(project)
}
