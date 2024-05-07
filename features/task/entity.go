package task

import "time"

type TaskEntity struct {
	ID          uint
	ProjectID   uint
	TaskName    string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DataTaskInterface interface {
	Insert(task TaskEntity) error
	SelectByProjectId(projectid uint) ([]TaskEntity, error)
	SelectById(id uint) (*TaskEntity, error)
	Delete(idtask uint) error
	Update(idtask uint, task TaskEntity) error
}

type ServiceTaskInterface interface {
	Create(task TaskEntity) error
	Delete(idtask uint) error
	Update(idtask uint, projectid uint, task TaskEntity) error
}
