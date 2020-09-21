package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectcorp/domain/model"
	projectPorts "projectcorp/ports/input/project"
	projectErrors "projectcorp/utils/errors"
)

func AddParticipantToProjectFactory(useCase projectPorts.IAddParticipantToProject) gin.HandlerFunc {
	addParticipantToProject := func(c *gin.Context){
		projectId := c.Param("id")
		var participant model.Participant
		participant.Project_id = projectId
		if errMarshal := c.Bind(&participant); errMarshal != nil {
			fmt.Println("error unmarshiling :", errMarshal)
			c.JSON(http.StatusBadRequest, errMarshal)
			return
		}

		err := useCase.AddParticipant(&participant)

		if err != nil {
			if err.Error() == "project owner is not working at the same department as the participant" {
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
	return addParticipantToProject
}
