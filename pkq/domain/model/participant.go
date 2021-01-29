package model

import "projectcorp/pkq/utils/enums"

type Participant struct {
	Id         string `json:"id"`
	Role       enums.Role
	Department string
	Project_id string
}

