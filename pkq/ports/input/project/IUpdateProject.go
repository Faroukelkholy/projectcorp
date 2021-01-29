package project

import "projectcorp/pkq/domain/model"

type IUpdateProject interface {
	UpdateProject(project *model.Project) error
}
