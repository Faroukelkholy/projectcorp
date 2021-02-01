package entity

import "projectcorp/pkq/utils/enums"

type Participant struct {
	tableName  struct{} `pg:"participant"`
	ID         string
	Role       enums.Role
	Department string
	ProjectID  string
}
