package server

import (
	"projectcorp/pkq/adapter/http/server/controllers"
	"projectcorp/pkq/port/input"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin      *gin.Engine
	usecases input.IDomainUseCases
}

func New(usecases input.IDomainUseCases) *Server {
	s := new(Server)
	s.gin = gin.Default()
	s.usecases = usecases
	s.initControllers()
	return s
}

func (s *Server) Start(port string) error {
	return s.gin.Run(port)
}

func (s *Server) initControllers() {
	getDevicesController := controllers.GetProjectsHandler(s.usecases)
	createDeviceController := controllers.CreateProjectHandler(s.usecases)
	updateDeviceController := controllers.UpdateProjectHandler(s.usecases)
	addParitcipantToProjectController := controllers.AddParticipantToProjectHandler(s.usecases)

	s.gin.GET("/projects", getDevicesController)
	s.gin.POST("/projects", createDeviceController)
	s.gin.PUT("/projects/:id", updateDeviceController)
	s.gin.POST("/projects/:id/participants", addParitcipantToProjectController)
}
