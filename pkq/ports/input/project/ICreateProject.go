package project

import "projectcorp/pkq/domain/model"

type ICreateProject interface {
	CreateProject(project *model.Project) error
}
