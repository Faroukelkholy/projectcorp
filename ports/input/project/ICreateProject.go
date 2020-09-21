package project

import "projectcorp/domain/model"

type ICreateProject interface {
	CreateProject(project *model.Project) error
}
