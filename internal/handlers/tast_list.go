package handlers

import (
	"TaskManager/internal/domain/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// createTask, getTasks, getTaskByID, updateTask, and deleteTask are HTTP handlers for managing tasks.
func (h *Handler) createTask(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "createTask"),
	)

	log.Info("getting user ID for task creation")
	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}
	log.Info("user ID retrieved successfully", zap.String("user_id", userId))

	var input model.TaskList
	if err := e.Bind(&input); err != nil {
		newErrorResponse(e, log, http.StatusBadRequest, err.Error())
		return nil
	}
	log.Info("binding input for task creation", zap.Any("input", input))
	id, err := h.services.TaskList.Create(userId, input)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}
	log.Info("task created successfully", zap.Int("task_id", id))

	return e.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// getAllTasksResponse is the response structure for retrieving all tasks.
type getAllTasksResponse struct {
	Data []model.TaskList `json:"data"`
}

// getTask retrieves a list of tasks.
func (h *Handler) getTasks(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "getTasks"),
	)

	log.Info("getting user ID for task retrieval")
	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}
	log.Info("user ID retrieved successfully", zap.String("user_id", userId))

	tasks, err := h.services.TaskList.GetAll(userId)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	log.Info("tasks retrieved successfully", zap.Int("task_count", len(tasks)))

	return e.JSON(http.StatusOK, getAllTasksResponse{
		Data: tasks,
	})
}

// getTaskByID retrieves a specific task by its ID.
func (h *Handler) getTaskByID(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "getTaskByID"),
	)

	log.Info("getting user ID for task retrieval by ID")
	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}
	log.Info("user ID retrieved successfully", zap.String("user_id", userId))

	taskId, err := strconv.Atoi(e.Param("id"))
	if taskId == 0 {
		newErrorResponse(e, log, http.StatusBadRequest, "task ID is required")
		return nil
	}
	log.Info("task ID retrieved successfully", zap.Int("task_id", taskId))

	task, err := h.services.TaskList.GetById(userId, taskId)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}
	log.Info("task retrieved successfully", zap.Any("task", task))

	return e.JSON(http.StatusOK, task)
}

// updateTask updates an existing task.
func (h *Handler) updateTask(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "updateTask"),
	)
	log.Info("getting user ID for task update")
	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}
	log.Info("user ID retrieved successfully", zap.String("user_id", userId))

	log.Info("getting task ID for update")
	taskId, err := strconv.Atoi(e.Param("id"))
	if taskId == 0 {
		newErrorResponse(e, log, http.StatusBadRequest, "task ID is required")
		return nil
	}
	log.Info("task ID retrieved successfully", zap.Int("task_id", taskId))

	var input model.UpdateTaskListInput
	if err := e.Bind(&input); err != nil {
		newErrorResponse(e, log, http.StatusBadRequest, err.Error())
		return nil
	}
	log.Info("binding input for task update", zap.Any("input", input))
	err = h.services.TaskList.Update(userId, taskId, input)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}
	log.Info("task updated successfully", zap.Int("task_id", taskId))
	return e.JSON(http.StatusOK, statusResponse{
		Status: "Task updated successfully",
	})
}

// deleteTask deletes a task by its ID.
func (h *Handler) deleteTask(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "createTask"),
	)

	log.Info("getting user ID for task deletion")
	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}

	log.Info("user ID retrieved successfully", zap.String("user_id", userId))
	taskId, err := strconv.Atoi(e.Param("id"))
	if taskId == 0 {
		newErrorResponse(e, log, http.StatusBadRequest, "task ID is required")
		return nil
	}
	log.Info("task ID retrieved successfully", zap.Int("task_id", taskId))
	err = h.services.TaskList.Delete(userId, taskId)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	log.Info("task deleted successfully", zap.Int("task_id", taskId))
	return e.JSON(http.StatusOK, statusResponse{
		Status: "Task deleted successfully",
	})

}
