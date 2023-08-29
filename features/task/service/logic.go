package service

import (
	"errors"
	"mia/my_task_app/features/task"
)

type TaskService struct {
	taskRepo task.TaskDataInterface
}

func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &TaskService{
		taskRepo: repo,
	}
}

// CreateTask implements task.TaskServiceInterface.
func (s *TaskService) Create(inputTask task.CoreTask, userID uint) (uint, error) {
	// Cek apakah user sudah login atau belum
	if userID == 0 {
		return 0, errors.New("user not logged in")
	}

	taskID, err := s.taskRepo.Insert(inputTask, userID)
	if err != nil {
		return 0, err
	}
	return taskID, nil

}

// GetAllTask implements task.TaskServiceInterface.
func (s *TaskService) GetAll(userID uint) ([]task.CoreTask, error) {
	result, err := s.taskRepo.SelectAll(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskById implements task.TaskServiceInterface.
func (s *TaskService) GetById(taskId uint, userID uint) (task.CoreTask, error) {
	result, err := s.taskRepo.Select(taskId, userID)
	if err != nil {
		return task.CoreTask{}, err
	}
	return result, nil
}

// UpdateTaskById implements task.TaskServiceInterface.
func (s *TaskService) UpdateById(taskId uint, userID uint, taskData task.CoreTask) error {
	err := s.taskRepo.Update(taskId, userID, taskData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTaskById implements task.TaskServiceInterface.
func (s *TaskService) DeleteById(taskId uint, userID uint) error {
	err := s.taskRepo.Delete(taskId, userID)
	if err != nil {
		return err
	}
	return nil
}
