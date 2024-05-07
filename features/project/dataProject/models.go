package dataProject

import (
	"my-task-app/features/task/dataTask"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	// ID        string `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID      uint
	ProjectName string
	Description string
	Tasks       []dataTask.Task `gorm:"foreignKey:ProjectID"`
}
