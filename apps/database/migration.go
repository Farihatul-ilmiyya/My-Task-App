package database

import (
	_projectData "mia/my_task_app/features/project/data"
	_taskData "mia/my_task_app/features/task/data"
	_userData "mia/my_task_app/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_projectData.Project{})
	db.AutoMigrate(&_taskData.Task{})
}
