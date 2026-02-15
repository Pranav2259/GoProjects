package storage

import "time"

type Status int

const (
	Pending Status = iota
	Completed
)

type Task struct {
	ID        int
	Task      string
	Status    Status
	CreatedAt time.Time
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Completed:
		return "completed"
	default:
		return "unknown"
	}
}
