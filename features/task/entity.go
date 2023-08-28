package task

import "time"

type CoreTask struct {
	ID               uint
	Title            string
	ProjectID        uint
	CompletionStatus string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
type TaskDataInterface interface {
	Insert(input CoreTask) (uint, error)
	SelectAll() ([]CoreTask, error)
	Select(taskId uint) (CoreTask, error)
	Update(taskId uint, taskData CoreTask) error
	Delete(taskId uint) error
}

type TaskServiceInterface interface {
	Create(input CoreTask) (uint, error)
	GetAll() ([]CoreTask, error)
	GetById(taskId uint) (CoreTask, error)
	UpdateById(taskId uint, taskData CoreTask) error
	DeleteById(taskId uint) error
}
