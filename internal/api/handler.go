package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flames31/jobqueue/internal/model"
	"github.com/flames31/jobqueue/internal/queue"
	"github.com/flames31/jobqueue/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	JobService *service.JobService
	JobQueue   *queue.JobQueue
}

func NewHandler(jobService *service.JobService, jobQueue *queue.JobQueue) *handler {
	return &handler{JobService: jobService, JobQueue: jobQueue}
}

func (h *handler) GETJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	job, err := h.JobService.ListJob(id)
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *handler) GETAllJobs(c *gin.Context) {

	jobs, err := h.JobService.ListAllJobs()
	if err != nil {
		log.Printf("Error getting all job : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *handler) POSTJob(c *gin.Context) {
	var req model.Job
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Result = "Not completed"
	req.Status = "todo"

	err := h.JobService.CreateJob(&req)
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JobQueue.Jobs <- req

	c.JSON(http.StatusOK, req)
}
