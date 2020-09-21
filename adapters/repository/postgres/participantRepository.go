package postgres

import (
	"errors"
	"fmt"
	"log"
	"projectcorp/adapters/repository/postgres/entity"
	"projectcorp/domain/model"
	"github.com/go-pg/pg/v9"
)

type ParticipantRepository struct {
	Database *pg.DB
}



func (thisPR *ParticipantRepository) CreateParticipant(participant *model.Participant) error{
	participantToSave := &entity.Participant{
		Id:         participant.Id,
		Role:       participant.Role,
		Department: participant.Department,
		Project_id: participant.Project_id,
	}
	result, err := thisPR.Database.Model(participantToSave).OnConflict("DO NOTHING").Insert()
	if err != nil {
		log.Println("error creating participant",err)
		return err
	}
	if result.RowsAffected() > 0 {
		fmt.Println("created")
		return nil

	} else {
		fmt.Println("did nothing")
		return errors.New("There a participant with this id")
	}
}