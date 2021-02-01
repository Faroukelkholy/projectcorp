package model

import "projectcorp/pkq/utils/enums"

type Project struct {
	ID         string `json:"id,omitempty"`
	Name       string
	State      enums.State `json:"state,omitempty"`
	Progress   string      `json:"progress,omitempty"`
	Department string
	Owner      string
}
