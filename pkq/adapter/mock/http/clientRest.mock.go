package mock

import "projectcorp/pkq/domain/model"

type Client struct {
	GetEmployeeMock func(idParam string) (*model.Employee, error)
}

func (cm *Client) GetEmployee(idParam string) (*model.Employee, error) {
	return cm.GetEmployeeMock(idParam)
}
