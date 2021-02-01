package controllers

import (
	"fmt"
	"net/http"
	projectports "projectcorp/pkq/port/input/project"

	"github.com/gin-gonic/gin"
)

func GetProjectsHandler(useCase projectports.IGetProjects) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
