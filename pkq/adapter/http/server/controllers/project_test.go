package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	httpmock "projectcorp/pkq/adapter/mock/http"
	repositorymock "projectcorp/pkq/adapter/mock/repository"
	mock "projectcorp/pkq/adapter/mock/use_case"
	model "projectcorp/pkq/domain/model"
	usecase "projectcorp/pkq/domain/use_case"
	input "projectcorp/pkq/port/input/project"
	"projectcorp/pkq/port/output"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	project = &model.Project{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	}
)

func setupUnitTestGetChemicals(ucMock input.ICreateProject, JSON []byte) (gin.HandlerFunc, *gin.Context, *httptest.ResponseRecorder) {
	//setup controller
	controller := CreateProjectHandler(ucMock)

	//setup echo context
	request, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(JSON))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = request

	return controller, c, rec
}

func TestCreateProjectSuccess(t *testing.T) {
	//init usecase mock
	mockFunc := func(project *model.Project) error {
		return nil
	}
	ucMock := &mock.CreateProjectUseCaseMock{CreateProjectMock: mockFunc}
	JSON, _ := json.Marshal(project)

	controller, c, rec := setupUnitTestGetChemicals(ucMock, JSON)

	controller(c)
	//assertions
	assert.EqualValues(t, http.StatusCreated, rec.Code)
}

func TestCreateProjectError(t *testing.T) {
	//init usecase mock
	mockFunc := func(project *model.Project) error {
		return nil
	}
	ucMock := &mock.CreateProjectUseCaseMock{CreateProjectMock: mockFunc}
	bodyString := `{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	}`
	JSON, _ := json.Marshal(bodyString)

	controller, c, rec := setupUnitTestGetChemicals(ucMock, JSON)
	controller(c)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
}

func setupIntTestCreateProject(rMock output.IProjectRepository, httpMock output.IHTTPClient, JSON []byte) (gin.HandlerFunc, *gin.Context, *httptest.ResponseRecorder) {
	//setup use case
	useCase := usecase.NewCreateProject(rMock, httpMock)

	//setup controller
	controller := CreateProjectHandler(useCase)

	//setup echo context
	request, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(JSON))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = request

	return controller, c, rec
}

func TestCreateProjectSuccessIntegration(t *testing.T) {
	//init repo mock
	rMockfunc := func(Project *model.Project) error {
		return nil
	}

	hcMockFunc := func(idParam string) (*model.Employee, error) {
		return &model.Employee{Role: "manager"}, nil
	}
	restRepoMock := &httpmock.Client{GetEmployeeMock: hcMockFunc}
	projectRepoMock := &repositorymock.ProjectRepo{CreateProjectMock: rMockfunc}

	JSON, _ := json.Marshal(project)

	controller, c, rec := setupIntTestCreateProject(projectRepoMock, restRepoMock, JSON)

	controller(c)
	//assertions
	assert.EqualValues(t, http.StatusCreated, rec.Code)
}

func TestCreateProjectErrorIntegration(t *testing.T) {
	//init repo mock
	rMockfunc := func(Project *model.Project) error {
		return nil
	}
	hcMock := func(idParam string) (*model.Employee, error) {
		return &model.Employee{Role: "employee"}, nil
	}
	restRepoMock := &httpmock.Client{GetEmployeeMock: hcMock}
	projectRepoMock := &repositorymock.ProjectRepo{CreateProjectMock: rMockfunc}

	bodyString := `{
		Name:       "testingProject",
		State:      "planned",
		Department: "sales",
		Owner:      "a0d5e87a-af04-473d-b1f5-3105bbf986c8",
	}`
	JSON, _ := json.Marshal(bodyString)

	controller, c, rec := setupIntTestCreateProject(projectRepoMock, restRepoMock, JSON)
	controller(c)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
}
