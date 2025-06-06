package main

import (
	"log"

	"github.com/flames31/jobqueue/internal/api"
	"github.com/flames31/jobqueue/internal/db"
	"github.com/flames31/jobqueue/internal/queue"
	"github.com/flames31/jobqueue/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERROR loading env : %v", err)
	}

	jobsDB, err := db.InitDB()
	if err != nil {
		log.Fatalf("ERROR during DB init : %v", err)
	}

	jobService := service.NewJobService(jobsDB)
	queue := queue.NewJobQueue(100, jobsDB)
	queue.Start(5)

	handler := api.NewHandler(jobService, queue)

	r := gin.Default()

	r.GET("/jobs", handler.GETAllJobs)
	r.GET("/jobs/:id", handler.GETJob)
	r.POST("/jobs", handler.POSTJob)
	r.Run()
}
