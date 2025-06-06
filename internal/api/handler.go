package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flames31/jobqueue/internal/model"
	"github.com/flames31/jobqueue/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	JobService *service.JobService
}

func NewHandler(jobService *service.JobService) *handler {
	return &handler{JobService: jobService}
}

func (h *handler) GETJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	job, err := h.JobService.ListJob(id)
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusNoContent, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, job)
}

func (h *handler) GETAllJobs(c *gin.Context) {

	jobs, err := h.JobService.ListAllJobs()
	if err != nil {
		log.Printf("Error getting all job : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *handler) POSTJob(c *gin.Context) {
	var req model.Job
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	err := h.JobService.CreateJob(&req)
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusNoContent, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job created succesfully",
	})
}
