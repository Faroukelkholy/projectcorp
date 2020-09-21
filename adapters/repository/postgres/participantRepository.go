package postgres

import (
	"github.com/go-pg/pg/v9"
	"log"
	"projectcorp/adapters/repository/postgres/entity"
	"projectcorp/domain/model"
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
	_, err := thisPR.Database.Model(participantToSave).OnConflict("DO NOTHING").Insert()
	if err != nil {
		log.Println("error creating participant",err)
		return err
	}
	return nil
}