package project

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"projectcorp/pkq/adapters/mock/clientRest"
	"projectcorp/pkq/adapters/mock/repository"
	"projectcorp/pkq/adapters/repository/postgres/entity"
	"projectcorp/pkq/domain/model"
	"projectcorp/pkq/domain/useCases"
	"projectcorp/pkq/domain/useCases/project"
	projectPorts "projectcorp/pkq/ports/input/project"
	"projectcorp/pkq/ports/output"
	"testing"
)

type projectUseCases struct {
	projectPorts.IGetProjects
	projectPorts.ICreateProject
	projectPorts.IUpdateProject
	projectPorts.IAddParticipantToProject
}

type databaseAdapter struct {
	output.IProjectRepository
	output.IParticipantRepository
}

func TestCreateProjectSuccessController(t *testing.T) {
	//init repo mock
	createProjectRepoMockfunc := func(Project *model.Project) error {
		return nil
	}
	GetEmployeeMock := func(idParam string) (*model.Employee,error){
		return &model.Employee{Role: "manager"},nil
	}
	restRepoMock := &clientRest.ClientRestMock{GetEmployeeMock}
	projectRepoMock := &repository.ProjectRepoMock{nil, nil,createProjectRepoMockfunc,nil}

	//init use Case
	createProjectUseCase := project.NewCreateProjectUseCase(projectRepoMock,restRepoMock)
	projectUseCase := &projectUseCases{nil,createProjectUseCase,nil,nil}
	var domainUseCase useCases.DomainUseCases
	domainUseCase.IProjectUseCases = projectUseCase

	project :=&entity.Project{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	}
	jsonValue, _ := json.Marshal(project)
	request, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = request

	//init controller
	createProjectController := CreateProjectFactory(domainUseCase)
	createProjectController(c)

	assert.EqualValues(t,http.StatusCreated,response.Code)

}

func TestCreateProjectErrorController(t *testing.T) {
	//init repo mock
	createProjectRepoMockfunc := func(Project *model.Project) error {
		return nil
	}
	GetEmployeeMock := func(idParam string) (*model.Employee,error){
		return &model.Employee{Role: "employee"},nil
	}
	restRepoMock := &clientRest.ClientRestMock{GetEmployeeMock}
	projectRepoMock := &repository.ProjectRepoMock{nil, nil,createProjectRepoMockfunc,nil}

	//init use Case
	createProjectUseCase := project.NewCreateProjectUseCase(projectRepoMock,restRepoMock)
	projectUseCase := &projectUseCases{nil,createProjectUseCase,nil,nil}
	var domainUseCase useCases.DomainUseCases
	domainUseCase.IProjectUseCases = projectUseCase

	bodyString :=`{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	}`
	jsonValue, _ := json.Marshal(bodyString)

	request, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = request

	//init controller
	createProjectController := CreateProjectFactory(domainUseCase)
	createProjectController(c)

	assert.EqualValues(t,http.StatusBadRequest,response.Code)

}
