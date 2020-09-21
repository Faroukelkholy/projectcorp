package gin

import (
	"github.com/gin-gonic/gin"
	"projectcorp/adapters/rest/controllers/project"
	"projectcorp/domain"
	"projectcorp/ports/input"
)

type GinServer struct {
	router *gin.Engine
}

func NewGinServer() input.IRestAdapterPort {
	ginServer := new(GinServer)
	ginServer.router = gin.Default()
	ginServer.initControllers()
	return ginServer
}

func (thisGinServer *GinServer) Start(port string) error{
	return thisGinServer.router.Run(port)
}

func (thisGinServer *GinServer) initControllers() {
	domainUseCaseSingleton :=domain.NewDomainUseCases()

	getDevicesController :=project.GetProjectsFactory(domainUseCaseSingleton.DomainUseCases.IProjectUseCases)
	createDeviceController :=project.CreateProjectFactory(domainUseCaseSingleton.DomainUseCases.IProjectUseCases)
	updateDeviceController := project.UpdateProjectFactory(domainUseCaseSingleton.DomainUseCases.IProjectUseCases)
	addParitcipantToProjectController := project.AddParticipantToProjectFactory(domainUseCaseSingleton.DomainUseCases.IProjectUseCases)

	thisGinServer.router.GET("/projects",getDevicesController)
	thisGinServer.router.POST("/projects",createDeviceController)
	thisGinServer.router.PUT("/projects/:id",updateDeviceController)
	thisGinServer.router.POST("/projects/:id/participants",addParitcipantToProjectController)

}