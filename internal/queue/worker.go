package queue

import (
	"fmt"
	"time"

	"github.com/flames31/jobqueue/internal/model"
	"gorm.io/gorm"
)

func startJob(id int, jobs <-chan model.Job, db *gorm.DB) {
	for job := range jobs {
		fmt.Printf("Starting Job %v on worker %v...\n", id, job.ID)
		db.Model(&job).Update("status", "in-progress")

		time.Sleep(15 * time.Second)

		job.Status = "done"
		job.Result = "Processed succesfully!"

		if db.Save(&job).Error != nil {
			fmt.Println("error saving job")
		}
		fmt.Printf("Job %v is done\n", id)
	}
}
