package postgres

import (
	"errors"
	"github.com/go-pg/pg/v9"
	"log"
	"projectcorp/adapters/repository/postgres/entity"
	"projectcorp/domain/model"
)

type ProjectRepository struct {
	Database *pg.DB
}



func (thisPR *ProjectRepository) GetProjects() ([]*model.Project,error){
	projectsQueried := []entity.Project{}
	err := thisPR.Database.Model(&projectsQueried).Select()

	if err != nil {
		return nil,err
	}

	var projects []*model.Project
	for _, project := range projectsQueried {
		result := &model.Project{
			Id:         project.Id,
			Name:       project.Name,
			State:      project.State,
			Progress:   project.Progress,
			Department: project.Department,
			Owner:      project.Owner,
		}
		projects = append(projects, result)
	}
	return projects, nil
}

func (thisPR *ProjectRepository) GetProject(id string) (*model.Project,error){
	projectQueried := &entity.Project{}
	err := thisPR.Database.Model(projectQueried).Where("id=?",id).First()

	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, errors.New("project does not exists")
		}
		return nil,err
	}

	project := &model.Project{
		Id:         projectQueried.Id,
		Name:       projectQueried.Name,
		State:      projectQueried.State,
		Progress:   projectQueried.Progress,
		Department: projectQueried.Department,
		Owner:      projectQueried.Owner,
	}
	return project, nil
}

func (thisPR *ProjectRepository) CreateProject(project *model.Project) error{
	projectToSave := &entity.Project{
		Name:       project.Name,
		State:      project.State,
		Progress:   project.Progress,
		Department: project.Department,
		Owner:      project.Owner,
	}
	_, err := thisPR.Database.Model(projectToSave).OnConflict("DO NOTHING").Insert()
	if err != nil {
		log.Println("error creating project",err)
		return err
	}
	return nil
}


func (thisPR *ProjectRepository) UpdateProject(project *model.Project) error{
	projectToUpdate := &entity.Project{
		Name:       project.Name,
		State:      project.State,
		Progress:   project.Progress,
		Department: project.Department,
		Owner:      project.Owner,
	}
	_, err := thisPR.Database.Model(projectToUpdate).Where("id=?",project.Id).Update(projectToUpdate)
	if err != nil {
		log.Println("Error in repo UpdateProject:",err)
		return err
	}
	return nil
}