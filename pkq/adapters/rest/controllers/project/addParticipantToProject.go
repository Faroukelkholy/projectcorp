package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectcorp/pkq/domain/model"
	projectPorts "projectcorp/pkq/ports/input/project"
	projectErrors "projectcorp/pkq/utils/errors"
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

		if participant.Id == "" || participant.Role == "" || participant.Department == "" {
			restErr := projectErrors.BadRequest("missing params that are required, required params are id && role && department")
			c.JSON(restErr.Status, restErr)
			return
		}

		err := useCase.AddParticipant(&participant)

		if err != nil {
			if err.Error() == "project owner is not working at the same department as the participant" || err.Error() =="no project exists for this id" {
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

