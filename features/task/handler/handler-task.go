package handler

import (
	"my-task-app/features/task"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.ServiceTaskInterface
}

func New(tk task.ServiceTaskInterface) *TaskHandler {
	return &TaskHandler{
		taskService: tk,
	}
}

func (th *TaskHandler) CreateTask(c echo.Context) error {
	// Membaca data dari request body
	var newTask TaskRequest
	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to parse request body: " + err.Error(),
		})
	}

	// Mapping dari request ke TaskEntity
	inputTask := task.TaskEntity{
		ProjectID:   newTask.IDProject,
		TaskName:    newTask.TaskName,
		Description: newTask.Description,
		Status:      newTask.Status,
	}

	// Memanggil method service layer
	if err := th.taskService.Create(inputTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to create task: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "task created successfully",
	})
}

func (th *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error converting id: " + errConv.Error(),
		})
	}

	var updateData TaskRequest
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error binding task: " + err.Error(),
		})
	}

	inputTask := task.TaskEntity{
		ProjectID:   updateData.IDProject,
		TaskName:    updateData.TaskName,
		Description: updateData.Description,
		Status:      updateData.Status,
	}

	if err := th.taskService.Update(uint(idConv), updateData.IDProject, inputTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error updating task: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "successfully updated task",
	})
}

func (th *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}
	err := th.taskService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete task " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete task",
	})
}
