package postgres

import (
	"errors"
	"log"
	"projectcorp/pkq/adapter/repository/postgres/entity"
	"projectcorp/pkq/domain/model"

	"github.com/go-pg/pg/v9"
)

type ProjectRepository struct {
	DB *pg.DB
}

func (pr *ProjectRepository) GetProjects() ([]*model.Project, error) {
	projectsQueried := []entity.Project{}
	err := pr.DB.Model(&projectsQueried).Select()

	if err != nil {
		return nil, err
	}

	var projects []*model.Project
	for _, project := range projectsQueried {
		result := &model.Project{
			ID:         project.ID,
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

func (pr *ProjectRepository) GetProject(id string) (*model.Project, error) {
	projectQueried := &entity.Project{}
	err := pr.DB.Model(projectQueried).Where("id=?", id).First()

	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, errors.New("project does not exists")
		}
		return nil, err
	}

	project := &model.Project{
		ID:         projectQueried.ID,
		Name:       projectQueried.Name,
		State:      projectQueried.State,
		Progress:   projectQueried.Progress,
		Department: projectQueried.Department,
		Owner:      projectQueried.Owner,
	}
	return project, nil
}

func (pr *ProjectRepository) CreateProject(project *model.Project) error {
	projectToSave := &entity.Project{
		Name:       project.Name,
		State:      project.State,
		Progress:   project.Progress,
		Department: project.Department,
		Owner:      project.Owner,
	}
	_, err := pr.DB.Model(projectToSave).OnConflict("DO NOTHING").Insert()
	if err != nil {
		log.Println("error creating project", err)
		return err
	}
	return nil
}

func (pr *ProjectRepository) UpdateProject(project *model.Project) error {
	projectToUpdate := &entity.Project{
		Name:       project.Name,
		State:      project.State,
		Progress:   project.Progress,
		Department: project.Department,
		Owner:      project.Owner,
	}
	_, err := pr.DB.Model(projectToUpdate).Where("id=?", project.ID).Update(projectToUpdate)
	if err != nil {
		log.Println("Error in repo UpdateProject:", err)
		return err
	}
	return nil
}
