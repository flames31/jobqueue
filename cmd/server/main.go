package main

import (
	"log"
	"os"

	"github.com/flames31/jobqueue/internal/api"
	"github.com/flames31/jobqueue/internal/db"
	"github.com/flames31/jobqueue/internal/middleware"
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

	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("ERROR during DB init : %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("NO JWT SECRET SET!")
	}
	newService := service.NewService(dbConn)
	queue := queue.NewJobQueue(100, dbConn)
	queue.Start(5)

	handler := api.NewHandler(newService, queue, jwtSecret)

	r := gin.Default()

	r.POST("/register", handler.POSTUserRegister)
	r.POST("/login", handler.POSTUserLogin)

	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware(jwtSecret))
	{
		protected.GET("/jobs", handler.GETAllJobs)
		protected.GET("/jobs/:id", handler.GETJob)
		protected.POST("/jobs", handler.POSTJob)
	}

	r.Run(":8080")
}
