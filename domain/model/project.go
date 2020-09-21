package model

import "projectcorp/utils/enums"

type Project struct {
	Id string
	Name string
	State enums.State
	Progress string
	Department string
	Owner string
}


