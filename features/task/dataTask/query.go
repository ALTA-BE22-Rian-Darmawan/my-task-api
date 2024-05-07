package dataTask

import (
	"my-task-app/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.DataTaskInterface {
	return &taskQuery{
		db: db,
	}
}

// Delete implements task.DataTaskInterface.
func (t *taskQuery) Delete(idtask uint) error {
	return t.db.Where("id = ?", idtask).Delete(&Task{}).Error
}

// Insert implements task.DataTaskInterface.
func (t *taskQuery) Insert(task task.TaskEntity) error {
	//var taskGorm Project
	taskGorm := Task{
		ProjectID:   task.ProjectID,
		TaskName:    task.TaskName,
		Description: task.Description,
		Status:      task.Status,
	}
	tx := t.db.Create(&taskGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectByProjectId implements task.DataTaskInterface.
func (t *taskQuery) SelectByProjectId(projectid uint) ([]task.TaskEntity, error) {
	var allTask []Task // var penampung data yg dibaca dari db
	tx := t.db.Where("project_id =?", projectid).Find(&allTask)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allTaskCore []task.TaskEntity
	for _, v := range allTask {
		allTaskCore = append(allTaskCore, task.TaskEntity{
			ID:          v.ID,
			ProjectID:   v.ProjectID,
			TaskName:    v.TaskName,
			Description: v.Description,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return allTaskCore, nil
}

// Update implements task.DataTaskInterface.
func (t *taskQuery) Update(idtask uint, task task.TaskEntity) error {
	//var taskGorm Project
	taskGorm := Task{
		ProjectID:   task.ProjectID,
		TaskName:    task.TaskName,
		Description: task.Description,
		Status:      task.Status,
	}
	tx := t.db.Model(&taskGorm).Where("id =?", idtask).Updates(&taskGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (t *taskQuery) SelectById(id uint) (*task.TaskEntity, error) {
	var taskGorm Task
	tx := t.db.First(&taskGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var projectcore = task.TaskEntity{
		ID:          taskGorm.ID,
		ProjectID:   taskGorm.ProjectID,
		TaskName:    taskGorm.TaskName,
		Description: taskGorm.Description,
		Status:      taskGorm.Status,
		CreatedAt:   taskGorm.CreatedAt,
		UpdatedAt:   taskGorm.UpdatedAt,
	}

	return &projectcore, nil
}
