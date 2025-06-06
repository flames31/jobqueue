package service

import (
	"fmt"

	"github.com/flames31/jobqueue/internal/model"
	"gorm.io/gorm"
)

type JobService struct {
	jobDB *gorm.DB
}

func NewJobService(jobDB *gorm.DB) *JobService {
	return &JobService{jobDB: jobDB}
}

func (js *JobService) CreateJob(job *model.Job) error {
	fmt.Println(*job)
	err := js.jobDB.Create(job).Error
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	return nil
}

func (js *JobService) ListAllJobs() ([]model.Job, error) {
	var jobs []model.Job
	err := js.jobDB.Find(&jobs).Error

	if err != nil {
		return []model.Job{}, fmt.Errorf("failed to list all jobs: %w", err)
	}

	return jobs, nil
}

func (js *JobService) ListJob(id int) (model.Job, error) {

	var job model.Job
	err := js.jobDB.First(&job, 1).Error

	if err != nil {
		return model.Job{}, fmt.Errorf("failed to list job: %w", err)
	}

	return job, nil
}
