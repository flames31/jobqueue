package queue

import "github.com/flames31/jobqueue/internal/model"

type JobQueue struct {
	Jobs chan model.Job
}

func NewJobQueue(bufferSize int) *JobQueue {
	return &JobQueue{
		Jobs: make(chan model.Job, bufferSize),
	}
}

func (q *JobQueue) Start(nodes int) {
	for i := 0; i < nodes; i++ {
		go startJob(i, q.Jobs)
	}
}
