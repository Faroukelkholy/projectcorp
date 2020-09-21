package project

type IProjectUseCases interface {
	IGetProjects
	ICreateProject
	IUpdateProject
	IAddParticipantToProject
}
