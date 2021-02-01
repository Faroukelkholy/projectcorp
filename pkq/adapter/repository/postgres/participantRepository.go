package postgres

import (
	"log"
	"projectcorp/pkq/adapter/repository/postgres/entity"
	"projectcorp/pkq/domain/model"

	"github.com/go-pg/pg/v9"
)

type ParticipantRepository struct {
	DB *pg.DB
}

func (pr *ParticipantRepository) CreateParticipant(participant *model.Participant) error {
	participantToSave := &entity.Participant{
		ID:         participant.ID,
		Role:       participant.Role,
		Department: participant.Department,
		ProjectID:  participant.ProjectID,
	}
	_, err := pr.DB.Model(participantToSave).OnConflict("DO NOTHING").Insert()
	if err != nil {
		log.Println("error creating participant", err)
		return err
	}
	return nil
}
