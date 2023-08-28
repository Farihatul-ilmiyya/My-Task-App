package service

import (
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
func (s *TaskService) Create(inputTask task.CoreTask) (uint, error) {

	taskID, err := s.taskRepo.Insert(inputTask)
	if err != nil {
		return 0, err
	}
	return taskID, nil

}

// GetAllTask implements task.TaskServiceInterface.
func (s *TaskService) GetAll() ([]task.CoreTask, error) {
	result, err := s.taskRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskById implements task.TaskServiceInterface.
func (s *TaskService) GetById(taskId uint) (task.CoreTask, error) {
	result, err := s.taskRepo.Select(taskId)
	if err != nil {
		return task.CoreTask{}, err
	}
	return result, nil
}

// UpdateTaskById implements task.TaskServiceInterface.
func (s *TaskService) UpdateById(taskId uint, taskData task.CoreTask) error {
	err := s.taskRepo.Update(taskId, taskData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTaskById implements task.TaskServiceInterface.
func (s *TaskService) DeleteById(taskId uint) error {
	err := s.taskRepo.Delete(taskId)
	if err != nil {
		return err
	}
	return nil
}
