package output

import "projectcorp/domain/model"

type IClientRest interface {
	GetEmployee(url string, idParam string) (*model.Employee,error)
}