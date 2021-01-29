package entity

import "projectcorp/pkq/utils/enums"

type Project struct {
	tableName  struct{}  `pg:"project"`
	Id         string
	Name       string
	State      enums.State
	Progress   string
	Department string
	Owner      string
}


