package project

import "projectcorp/pkq/domain/model"

type IGetProjects interface {
	GetProjects() ([]*model.Project,error)
}