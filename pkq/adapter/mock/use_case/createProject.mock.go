package mock

import "projectcorp/pkq/domain/model"

type CreateProjectUseCaseMock struct {
	CreateProjectMock func(project *model.Project) error
}

func (c *CreateProjectUseCaseMock) CreateProject(project *model.Project) error {
	return c.CreateProjectMock(project)
}
