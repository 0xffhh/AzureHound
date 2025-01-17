package models

type CompleteJobRequest struct {
	Status     string    `json:"status"`
	StatusEnum JobStatus `json:"-"`
	Message    string    `json:"message"`
}

type JobStatus int

const (
	JobStatusInvalid   JobStatus = -1
	JobStatusReady     JobStatus = 0
	JobStatusRunning   JobStatus = 1
	JobStatusComplete  JobStatus = 2
	JobStatusCanceled  JobStatus = 3
	JobStatusTimedOut  JobStatus = 4
	JobStatusFailed    JobStatus = 5
	JobStatusIngesting JobStatus = 6
)

func (s JobStatus) String() string {
	switch s {
	case JobStatusReady:
		return "READY"

	case JobStatusRunning:
		return "RUNNING"

	case JobStatusComplete:
		return "COMPLETE"

	case JobStatusCanceled:
		return "CANCELED"

	case JobStatusTimedOut:
		return "TIMEDOUT"

	case JobStatusFailed:
		return "FAILED"

	case JobStatusIngesting:
		return "INGESTING"

	default:
		return "INVALIDSTATUS"
	}
}
