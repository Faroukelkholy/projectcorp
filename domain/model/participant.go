package model

import "projectcorp/utils/enums"

type Participant struct {
	Id string
	Role enums.Role
	Department string
	Project_id string
}

