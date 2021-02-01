package controllers

import (
	"net/http"
	"projectcorp/pkq/domain/model"
	projectports "projectcorp/pkq/port/input/project"
	projecterrors "projectcorp/pkq/utils/errors"

	"github.com/gin-gonic/gin"
)

func AddParticipantToProjectHandler(useCase projectports.IAddParticipantToProject) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := c.Param("id")
		var participant model.Participant
		participant.ProjectID = projectID
		if errMarshal := c.Bind(&participant); errMarshal != nil {
			c.JSON(http.StatusBadRequest, errMarshal)
			return
		}

		if participant.ID == "" || participant.Role == "" || participant.Department == "" {
			restErr := projecterrors.BadRequest("missing params that are required, required params are id && role && department")
			c.JSON(restErr.Status, restErr)
			return
		}

		err := useCase.AddParticipant(&participant)

		if err != nil {
			if err.Error() == "project owner is not working at the same department as the participant" || err.Error() == "no project exists for this id" {
				restErr := projecterrors.BadRequest(err.Error())
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
}
