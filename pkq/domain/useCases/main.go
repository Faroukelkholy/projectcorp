package useCases

import (
	"projectcorp/pkq/domain/useCases/project"
	projectPorts "projectcorp/pkq/ports/input/project"
	"projectcorp/pkq/ports/output"
)

type ProjectUseCases struct {
	projectPorts.IGetProjects
	projectPorts.ICreateProject
	projectPorts.IUpdateProject
	projectPorts.IAddParticipantToProject
}
type DomainUseCases struct {
	databaseAdapter   output.IDatabasePort
	restClientAdapter output.IClientRest
	projectPorts.IProjectUseCases
}

func NewDomainUseCases(databaseAdapter output.IDatabasePort,restClientAdapter output.IClientRest) DomainUseCases {
	domainUseCase := DomainUseCases{}
	domainUseCase.databaseAdapter = databaseAdapter
	domainUseCase.restClientAdapter = restClientAdapter
	domainUseCase.initUseCases()
	return domainUseCase
}

func (thisDUC *DomainUseCases) initUseCases() {
	thisDUC.IProjectUseCases = &ProjectUseCases{
		project.NewGetProjectsUseCase(thisDUC.databaseAdapter),
		project.NewCreateProjectUseCase(thisDUC.databaseAdapter,thisDUC.restClientAdapter),
		project.NewUpdateProjectUseCase(thisDUC.databaseAdapter,thisDUC.restClientAdapter),
		project.NewAddParticipantToProjectUseCase(thisDUC.databaseAdapter,thisDUC.databaseAdapter),
	}
}