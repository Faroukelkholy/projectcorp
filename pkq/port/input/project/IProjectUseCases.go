package input

type IProjectUseCases interface {
	IGetProjects
	ICreateProject
	IUpdateProject
	IAddParticipantToProject
}
