package model

import "projectcorp/utils/enums"

type Employee struct {
	Id string
	FirstName string
	LastName string
	Email string
	Department string
	Role enums.Role
}
