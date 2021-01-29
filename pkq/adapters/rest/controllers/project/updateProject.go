package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/ports/input/project"
	projectErrors "projectcorp/pkq/utils/errors"
)

func UpdateProjectFactory(useCase projectPorts.IUpdateProject) gin.HandlerFunc {
	updateProject := func(c *gin.Context){
		projectId := c.Param("id")
		var project model.Project
		project.Id = projectId
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

		err := useCase.UpdateProject(&project)

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

		c.JSON(200, map[string]interface{}{
			"title": "Success",
		})
	}
	return updateProject
}