package api

import (
	"github.com/flames31/jobqueue/internal/pubsub"
	"github.com/flames31/jobqueue/internal/queue"
	"github.com/flames31/jobqueue/internal/service"
)

type handler struct {
	Service   *service.Service
	JobQueue  *queue.JobQueue
	JWTSecret string
	RDSP      pubsub.Publisher
}

func NewHandler(service *service.Service, jobQueue *queue.JobQueue, JwtSecret string, rdsP pubsub.Publisher) *handler {
	return &handler{Service: service, JobQueue: jobQueue, JWTSecret: JwtSecret, RDSP: rdsP}
}
