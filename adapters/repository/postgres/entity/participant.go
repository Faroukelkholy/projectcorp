package entity

import "projectcorp/utils/enums"

type Participant struct {
	tableName    struct{}  `pg:"participant"`
	Id string
	Role enums.Role
	Department string
	Project_id string
}
