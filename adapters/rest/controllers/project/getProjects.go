package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	projectPorts "projectcorp/ports/input/project"
)

func GetProjectsFactory(useCase projectPorts.IGetProjects) gin.HandlerFunc {
	getProjects := func(c *gin.Context){
		result, err := useCase.GetProjects()
		if err != nil {
			fmt.Println("getProjects UseCase", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(200, map[string]interface{}{
			"data": result,
		})
	}

	return getProjects
}
