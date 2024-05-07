package project

import (
	"my-task-app/features/task/dataTask"
	"time"
)

type ProjectEntity struct {
	ID          uint
	UserID      uint
	ProjectName string
	Description string
	Tasks       []dataTask.Task
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DataProjectInterface interface {
	Insert(project ProjectEntity) error
	Delete(id uint) error
	SelectById(id uint) (*ProjectEntity, error)
	Update(id uint, project ProjectEntity) error
	SelectByUserId(id uint) ([]ProjectEntity, error)
}

type ServiceProjectInterface interface {
	Create(project ProjectEntity) error
	Delete(id uint, userid uint) error
	Update(id uint, userid uint, project ProjectEntity) error
	GetById(id uint) (project *ProjectEntity, err error)
	GetByUserId(id uint) ([]ProjectEntity, error)
}
