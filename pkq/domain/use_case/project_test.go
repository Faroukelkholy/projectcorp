package usecase

import (
	httpmock "projectcorp/pkq/adapter/mock/http"
	repositorymock "projectcorp/pkq/adapter/mock/repository"
	"projectcorp/pkq/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProjectSuccess(t *testing.T) {
	//setup mocks
	rMockFunc := func(project *model.Project) error {
		return nil
	}
	hcMockFunc := func(idParam string) (*model.Employee, error) {
		return &model.Employee{Role: "manager"}, nil
	}
	hcMock := &httpmock.Client{GetEmployeeMock: hcMockFunc}
	rMock := &repositorymock.ProjectRepo{CreateProjectMock: rMockFunc}

	//execute service logic
	createProjectUseCase := NewCreateProject(rMock, hcMock)
	err := createProjectUseCase.CreateProject(&model.Project{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "bed8e365-3d7e-4994-960d-e51275343f23",
	})

	assert.Nil(t, err)
}

func TestCreateProjectError(t *testing.T) {
	//setup mocks
	rMockFunc := func(Project *model.Project) error {
		return nil
	}
	hcMockFunc := func(idParam string) (*model.Employee, error) {
		return &model.Employee{Role: "employee"}, nil
	}
	hcMock := &httpmock.Client{GetEmployeeMock: hcMockFunc}
	rMock := &repositorymock.ProjectRepo{CreateProjectMock: rMockFunc}

	//execute service logic
	createProjectUseCase := NewCreateProject(rMock, hcMock)
	err := createProjectUseCase.CreateProject(&model.Project{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	})
	assert.NotNil(t, err)
	assert.EqualValues(t, "project owner is not a manager", err.Error())
}
