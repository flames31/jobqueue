package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/flames31/jobqueue/internal/model"
	"github.com/flames31/jobqueue/internal/pubsub"
	"github.com/gin-gonic/gin"
)

func (h *handler) GETJob(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	job, err := h.Service.JobService.ListJob(id)
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")
	if userID == 0 {
		log.Println("Error : User does not exist")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User does not exist",
		})
		return
	}
	if job.UserID != userID {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Job not owned by user",
		})
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *handler) GETAllJobs(c *gin.Context) {

	userID := c.GetUint("user_id")
	if userID == 0 {
		log.Println("Error : User does not exist")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User does not exist",
		})
		return
	}
	jobs, err := h.Service.JobService.ListAllJobs(userID)
	if err != nil {
		log.Printf("Error getting all jobs : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *handler) POSTJob(c *gin.Context) {
	var job model.Job
	if err := c.ShouldBindBodyWithJSON(&job); err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")
	if userID == 0 {
		log.Println("Error : User does not exist")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User does not exist",
		})
		return
	}

	job.Result = "Not completed"
	job.Status = "todo"
	job.UserID = userID

	err := h.Service.JobService.CreateJob(&job)
	if err != nil {
		log.Printf("Error getting job : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.RDSP.Publish(pubsub.JobEvent{
		Type:    pubsub.EventJobCreated,
		JobID:   job.ID,
		UserID:  job.UserID,
		Payload: map[string]interface{}{"testing Job number": job.ID},
		Time:    time.Now(),
	})

	c.JSON(http.StatusOK, job)
}
