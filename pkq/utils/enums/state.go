package enums

type State string

const (
	ACTIVE  State = "active"
	PLANNED State = "planned"
	DONE    State = "done"
	FAILED  State = "failed"
)
