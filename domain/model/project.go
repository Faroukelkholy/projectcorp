package model

import "projectcorp/utils/enums"

type Project struct {
	Id string `json:”id,omitempty”`
	Name string
	State enums.State `json:”state,omitempty”`
	Progress string `json:”progress,omitempty”`
	Department string
	Owner string
}


