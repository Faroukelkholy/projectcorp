package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	projectErrors "projectcorp/utils/errors"
)

func CreateProjectFactory(useCase projectPorts.ICreateProject) gin.HandlerFunc {
	createProject := func(c *gin.Context){
		var project model.Project

		if errMarshal := c.Bind(&project); errMarshal != nil {
			fmt.Println("error unmarshiling :", errMarshal)
			c.JSON(http.StatusBadRequest, errMarshal)
			return
		}

		if project.Owner == "" || project.Name == "" || project.Department == "" {
			restErr := projectErrors.BadRequest("missing params that are required, required params are Owner && name && department")
			c.JSON(restErr.Status, restErr)
			return
		}

		err := useCase.CreateProject(&project)

		if err != nil {
			//if err.Error() == "\"project owner is not a manager\"" {
			if err.Error() == "project owner is not a manager" {
				restErr := projectErrors.BadRequest(err.Error())
				c.JSON(restErr.Status, restErr)
				return
			}
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(201, map[string]interface{}{
			"title": "Success",
		})
	}
	return createProject
}

