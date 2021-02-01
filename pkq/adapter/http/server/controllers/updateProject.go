package controllers

import (
	"fmt"
	"net/http"
	"projectcorp/pkq/domain/model"
	projectports "projectcorp/pkq/port/input/project"
	projecterrors "projectcorp/pkq/utils/errors"

	"github.com/gin-gonic/gin"
)

func UpdateProjectHandler(useCase projectports.IUpdateProject) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := c.Param("id")
		var project model.Project
		project.ID = projectID
		if errMarshal := c.Bind(&project); errMarshal != nil {
			fmt.Println("error unmarshiling :", errMarshal)
			c.JSON(http.StatusBadRequest, errMarshal)
			return
		}

		if project.Owner == "" || project.Name == "" || project.Department == "" {
			restErr := projecterrors.BadRequest("missing params that are required, required params are Owner && name && department")
			c.JSON(restErr.Status, restErr)
			return
		}

		err := useCase.UpdateProject(&project)

		if err != nil {
			//if err.Error() == "\"project owner is not a manager\"" {
			if err.Error() == "project owner is not a manager" {
				restErr := projecterrors.BadRequest(err.Error())
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
}
