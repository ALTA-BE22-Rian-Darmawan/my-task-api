package dataTask

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	// ID        string `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	ProjectID   uint
	TaskName    string
	Description string
	Status      string
}
