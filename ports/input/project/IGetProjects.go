package project

import "projectcorp/domain/model"

type IGetProjects interface {
	GetProjects() ([]*model.Project,error)
}