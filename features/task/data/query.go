package data

import (
	"errors"
	"mia/my_task_app/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: database,
	}
}

// Insert implements task.TaskDataInterface.
func (r *taskQuery) Insert(input task.CoreTask) (uint, error) {
	newTask := MapCoreTaskToTask(input)

	//simpan ke db
	tx := r.db.Create(&newTask)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}
	return newTask.ID, nil
}

// SelectAll implements task.TaskDataInterface.
func (r *taskQuery) SelectAll() ([]task.CoreTask, error) {
	var dataTask []Task
	tx := r.db.Find(&dataTask)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	//mapping Task Model to CoreTask
	coreTaskSlice := ListMapTaskToCoreTask(dataTask)
	return coreTaskSlice, nil
}

// Select implements task.TaskDataInterface.
func (r *taskQuery) Select(taskId uint) (task.CoreTask, error) {
	var taskData Task
	tx := r.db.First(&taskData, taskId)
	if tx.Error != nil {
		return task.CoreTask{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task.CoreTask{}, errors.New("data not found")
	}
	//Mapping Task to CorePtask
	coreTask := MapTaskToCoreTask(taskData)
	return coreTask, nil
}

// Update implements task.TaskDataInterface.
func (r *taskQuery) Update(taskId uint, taskData task.CoreTask) error {
	var task Task
	tx := r.db.First(&task, taskId)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}

	//Mapping Task to CoreTask
	updatedTask := MapCoreTaskToTask(taskData)
	tx = r.db.Model(&task).Updates(updatedTask)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}

// Delete implements task.TaskDataInterface.
func (r *taskQuery) Delete(taskId uint) error {
	var task Task
	tx := r.db.Delete(&task, taskId)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
