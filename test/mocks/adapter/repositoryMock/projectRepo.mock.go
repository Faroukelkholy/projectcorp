package repositoryMock

import "projectcorp/domain/model"

type ProjectRepoMock struct {
	GetProjectsMock func() ([]*model.Project,error)
	GetProjectMock func(id string) (*model.Project,error)
	CreateProjectMock func(project *model.Project) error
	UpdateProjectMock func(project *model.Project) error
}

func(thisPR *ProjectRepoMock) GetProjects() ([]*model.Project,error){
	return thisPR.GetProjectsMock()
}
func(thisPR *ProjectRepoMock) GetProject(id string) (*model.Project,error){
	return thisPR.GetProjectMock(id)
}

func(thisPR *ProjectRepoMock) CreateProject(project *model.Project) error{
	return thisPR.CreateProjectMock(project)
}

func(thisPR *ProjectRepoMock) UpdateProject(project *model.Project) error{
	return thisPR.UpdateProjectMock(project)
}





