package output

import "projectcorp/pkq/domain/model"

type IClientRest interface {
	GetEmployee(idParam string) (*model.Employee,error)
}