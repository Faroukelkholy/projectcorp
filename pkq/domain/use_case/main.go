package usecase

import (
	projectPorts "projectcorp/pkq/port/input/project"
	"projectcorp/pkq/port/output"
)

type ProjectUseCases struct {
	projectPorts.IGetProjects
	projectPorts.ICreateProject
	projectPorts.IUpdateProject
	projectPorts.IAddParticipantToProject
}
type DomainUseCases struct {
	DB         output.IRepositoryPort
	HTTPClient output.IHTTPClient
	projectPorts.IProjectUseCases
}

func New(DB output.IRepositoryPort, httpClient output.IHTTPClient) DomainUseCases {
	domainUseCase := DomainUseCases{}
	domainUseCase.DB = DB
	domainUseCase.HTTPClient = httpClient
	domainUseCase.initUseCases()
	return domainUseCase
}

func (d *DomainUseCases) initUseCases() {
	d.IProjectUseCases = &ProjectUseCases{
		NewGetProjects(d.DB),
		NewCreateProject(d.DB, d.HTTPClient),
		NewUpdateProject(d.DB, d.HTTPClient),
		NewAddParticipantToProject(d.DB, d.DB),
	}
}
