package model

import "projectcorp/pkq/utils/enums"

type Participant struct {
	ID         string `json:"id"`
	Role       enums.Role
	Department string
	ProjectID  string
}
