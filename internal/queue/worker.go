package queue

import (
	"fmt"
	"time"

	"github.com/flames31/jobqueue/internal/model"
)

func startJob(id int, jobs <-chan model.Job) {
	for job := range jobs {
		fmt.Printf("Starting Job %v on worker %v...\n", id, job.ID)

		time.Sleep(15 * time.Second)

		fmt.Printf("Job %v is done\n", id)
	}
}
