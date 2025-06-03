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

	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}

	var input model.TaskList
	if err := e.Bind(&input); err != nil {
		newErrorResponse(e, log, http.StatusBadRequest, err.Error())
		return nil
	}
	id, err := h.services.TaskList.Create(userId, input)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllTasksResponse struct {
	Data []model.TaskList `json:"data"`
}

// getTask retrieves a list of tasks.
func (h *Handler) getTasks(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "getTasks"),
	)

	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}

	tasks, err := h.services.TaskList.GetAll(userId)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	return e.JSON(http.StatusOK, getAllTasksResponse{
		Data: tasks,
	})
}

// getTaskByID retrieves a specific task by its ID.
func (h *Handler) getTaskByID(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "getTaskByID"),
	)

	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}

	taskId, err := strconv.Atoi(e.Param("id"))
	if taskId == 0 {
		newErrorResponse(e, log, http.StatusBadRequest, "task ID is required")
		return nil
	}

	task, err := h.services.TaskList.GetById(userId, taskId)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	return e.JSON(http.StatusOK, task)
}

// updateTask updates an existing task.
func (h *Handler) updateTask(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "updateTask"),
	)
	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}
	taskId, err := strconv.Atoi(e.Param("id"))
	if taskId == 0 {
		newErrorResponse(e, log, http.StatusBadRequest, "task ID is required")
		return nil
	}
	var input model.UpdateTaskListInput
	if err := e.Bind(&input); err != nil {
		newErrorResponse(e, log, http.StatusBadRequest, err.Error())
		return nil
	}
	err = h.services.TaskList.Update(userId, taskId, input)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}
	return e.JSON(http.StatusOK, statusResponse{
		Status: "Task updated successfully",
	})
}

// deleteTask deletes a task by its ID.
func (h *Handler) deleteTask(e echo.Context) error {
	log := h.logger.With(
		zap.String("handler", "createTask"),
	)

	userId, err := getUserId(e)
	if err != nil {
		newErrorResponse(e, log, http.StatusUnauthorized, err.Error())
		return nil
	}

	taskId, err := strconv.Atoi(e.Param("id"))
	if taskId == 0 {
		newErrorResponse(e, log, http.StatusBadRequest, "task ID is required")
		return nil
	}
	err = h.services.TaskList.Delete(userId, taskId)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	return e.JSON(http.StatusOK, statusResponse{
		Status: "Task deleted successfully",
	})

}
