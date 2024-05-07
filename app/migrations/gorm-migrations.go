package migrations

import (
	"my-task-app/features/project/dataProject"
	"my-task-app/features/task/dataTask"
	"my-task-app/features/user/data"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&data.User{})
	db.AutoMigrate(&dataProject.Project{})
	db.AutoMigrate(&dataTask.Task{})
}
