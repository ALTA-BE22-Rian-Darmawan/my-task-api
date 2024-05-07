package service

import (
	"errors"
	"my-task-app/features/task"
)

type taskService struct {
	taskData task.DataTaskInterface
}

func New(td task.DataTaskInterface) task.ServiceTaskInterface {
	return &taskService{
		taskData: td,
	}

}

// Create implements task.ServiceTaskInterface.
func (ts *taskService) Create(task task.TaskEntity) error {
	if task.TaskName == "" || task.ProjectID <= 0 {
		return errors.New("task name or project id cannot be empty")
	}

	if task.Status != "completed" && task.Status != "not completed" {
		return errors.New("task status must be completed or not completed")
	}

	err := ts.taskData.Insert(task)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements task.ServiceTaskInterface.
func (ts *taskService) Delete(idtask uint) error {
	if idtask <= 0 {
		return errors.New("invalid project ID")
	}

	return ts.taskData.Delete(idtask)
}

// Update implements task.ServiceTaskInterface.
func (ts *taskService) Update(idtask uint, projectid uint, task task.TaskEntity) error {
	if idtask <= 0 {
		return errors.New("invalid project ID")
	}

	if task.TaskName == "" || task.ProjectID <= 0 {
		return errors.New("task name or project id cannot be empty")
	}

	if task.Status != "completed" && task.Status != "not completed" {
		return errors.New("task status must be completed or not completed")
	}

	cekprojectid, err := ts.taskData.SelectById(idtask)
	if err != nil {
		return err
	}

	if cekprojectid.ProjectID != projectid {
		return errors.New("project id not match, cannot update task")
	}

	return ts.taskData.Update(idtask, task)
}
