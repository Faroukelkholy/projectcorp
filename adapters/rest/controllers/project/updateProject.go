package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	projectErrors "projectcorp/utils/errors"
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