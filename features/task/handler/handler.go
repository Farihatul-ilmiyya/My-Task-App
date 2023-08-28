package handler

import (
	"mia/my_task_app/apps/middlewares"
	"mia/my_task_app/features/task"
	"mia/my_task_app/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type taskHandler struct {
	taskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *taskHandler {
	return &taskHandler{
		taskService: service,
	}
}

func (h *taskHandler) CreateTask(c echo.Context) error {

	NewTask := new(TaskRequest)

	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&NewTask)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//mapping dari request to CoreTask
	input := MapTaskReqToCoreTask(*NewTask)
	_, err = h.taskService.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success create task", nil))
}

func (h *taskHandler) GetAllTask(c echo.Context) error {
	result, err := h.taskService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var tasksResponse []TaskResponse
	for _, v := range result {
		tasksResponse = append(tasksResponse, MapCoreTaskToTaskRes(v))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", tasksResponse))
}

func (h *taskHandler) GetTaskById(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.taskService.GetById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := TaskResponse{
		Title:     result.Title,
		ProjectID: result.ProjectID,
		CreatedAt: result.CreatedAt,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", resultResponse))
}

func (h *taskHandler) UpdateTaskById(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	taskInput := TaskRequest{}
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//Mapping task reques to core task
	Core := MapTaskReqToCoreTask(taskInput)
	err := h.taskService.UpdateById(uint(userID), Core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
	}

	// Get task data for response
	task, err := h.taskService.GetById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "task not found", nil))
	}
	resultResponse := MapCoreTaskToTaskRes(task)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "task updated successfully", resultResponse))
}

func (h *taskHandler) DeleteTaskById(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	err := h.taskService.DeleteById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete task", nil))
}
