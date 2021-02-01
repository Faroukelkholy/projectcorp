package output

import "projectcorp/pkq/domain/model"

type IHTTPClient interface {
	GetEmployee(idParam string) (*model.Employee, error)
}
