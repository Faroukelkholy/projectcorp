package project

import "projectcorp/domain/model"

type IUpdateProject interface {
	UpdateProject(project *model.Project) error
}
