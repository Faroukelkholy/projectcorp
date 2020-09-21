package clientRestMock

import "projectcorp/domain/model"

type ClientRestMock struct {
	GetEmployeeMock func(idParam string) (*model.Employee,error)
}

func(thisMock *ClientRestMock) GetEmployee(idParam string) (*model.Employee,error) {
	return thisMock.GetEmployeeMock(idParam)
}
