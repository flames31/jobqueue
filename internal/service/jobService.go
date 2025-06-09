package service

import (
	"fmt"

	"github.com/flames31/jobqueue/internal/model"
	"gorm.io/gorm"
)

type JobService struct {
	jobDB *gorm.DB
}

func (js *JobService) CreateJob(job *model.Job) error {
	err := js.jobDB.Create(job).Error
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	return nil
}

func (js *JobService) ListAllJobs(userID uint) ([]model.Job, error) {
	var jobs []model.Job
	err := js.jobDB.Model(&model.User{}).Association("Jobs").Find(&jobs)

	if err != nil {
		return []model.Job{}, fmt.Errorf("failed to list all jobs: %w", err)
	}

	return jobs, nil
}

func (js *JobService) ListJob(jobID int) (model.Job, error) {

	var job model.Job
	err := js.jobDB.First(&job, jobID).Error

	if err != nil {
		return model.Job{}, fmt.Errorf("failed to list job: %w", err)
	}

	return job, nil
}
