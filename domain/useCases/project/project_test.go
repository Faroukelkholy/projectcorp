package project

import (
	"github.com/stretchr/testify/assert"
	"log"
	"projectcorp/domain/model"
	"projectcorp/test/mocks/adapter/clientRestMock"
	"projectcorp/test/mocks/adapter/repositoryMock"
	"testing"
)

func TestCreateProjectSuccess(t *testing.T) {
	
	createProjectRepoMockfunc := func(project *model.Project) error {
		return nil
	}

	GetEmployeeMock := func(idParam string) (*model.Employee,error){
		return &model.Employee{Role: "manager"},nil
	}
	restRepoMock := &clientRestMock.ClientRestMock{GetEmployeeMock}
	projectRepoMock := &repositoryMock.ProjectRepoMock{nil, nil,createProjectRepoMockfunc,nil}
	createProjectUseCase := NewCreateProjectUseCase(projectRepoMock,restRepoMock)
	err := createProjectUseCase.CreateProject(&model.Project{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "bed8e365-3d7e-4994-960d-e51275343f23",
	})
	log.Println("err in  TestCreateProjectSuccess :", err)
	assert.Nil(t, err)
}

func TestCreateProjectError(t *testing.T) {

	createProjectRepoMockfunc := func(Project *model.Project) error {
		return nil
	}
	GetEmployeeMock := func(idParam string) (*model.Employee,error){
		 return &model.Employee{Role: "employee"},nil
	}
	restRepoMock := &clientRestMock.ClientRestMock{GetEmployeeMock}
	projectRepoMock := &repositoryMock.ProjectRepoMock{nil, nil,createProjectRepoMockfunc,nil}
	createProjectUseCase := NewCreateProjectUseCase(projectRepoMock,restRepoMock)
	err := createProjectUseCase.CreateProject(&model.Project{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	})
	assert.NotNil(t, err)
	assert.EqualValues(t, "project owner is not a manager", err.Error())
}