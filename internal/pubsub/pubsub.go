package pubsub

import "time"

type JobEvent struct {
	Type    string
	JobID   uint
	UserID  uint
	Payload map[string]interface{}
	Time    time.Time
}

const (
	EventJobCreated = "task.created"
	EventJobUpdated = "task.updated"
	EventJobDeleted = "task.deleted"
)

type Publisher interface {
	Publish(event JobEvent) error
}
