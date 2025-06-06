package queue

import (
	"github.com/flames31/jobqueue/internal/model"
	"gorm.io/gorm"
)

type JobQueue struct {
	Jobs chan model.Job
	DB   *gorm.DB
}

func NewJobQueue(bufferSize int, db *gorm.DB) *JobQueue {
	return &JobQueue{
		Jobs: make(chan model.Job, bufferSize),
		DB:   db,
	}
}

func (q *JobQueue) Start(nodes int) {
	for i := 0; i < nodes; i++ {
		go startJob(i, q.Jobs, q.DB)
	}
}
