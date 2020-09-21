package output

import "projectcorp/domain/model"

type IClientRest interface {
	GetEmployee(idParam string) (*model.Employee,error)
}