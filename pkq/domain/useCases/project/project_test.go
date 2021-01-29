package project

import (
	"github.com/stretchr/testify/assert"
	"log"
	"projectcorp/pkq/adapters/mock/clientRest"
	"projectcorp/pkq/adapters/mock/repository"
	"projectcorp/pkq/domain/model"
	"testing"
)

func TestCreateProjectSuccess(t *testing.T) {
	
	createProjectRepoMockfunc := func(project *model.Project) error {
		return nil
	}

	GetEmployeeMock := func(idParam string) (*model.Employee,error){
		return &model.Employee{Role: "manager"},nil
	}
	restRepoMock := &clientRest.ClientRestMock{GetEmployeeMock}
	projectRepoMock := &repository.ProjectRepoMock{nil, nil,createProjectRepoMockfunc,nil}
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
	restRepoMock := &clientRest.ClientRestMock{GetEmployeeMock}
	projectRepoMock := &repository.ProjectRepoMock{nil, nil,createProjectRepoMockfunc,nil}
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