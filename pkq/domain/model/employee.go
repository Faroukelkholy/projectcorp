package model

import "projectcorp/pkq/utils/enums"

type Employee struct {
	ID         string `json:"id,omitempty"`
	FirstName  string
	LastName   string
	Email      string
	Department string
	Role       enums.Role
}
