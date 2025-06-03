package service

import (
	"TaskManager/internal/domain/model"
	"TaskManager/internal/repository"
	"errors"
)

// TaskListService provides methods to manage task lists for users.
type TaskListService struct {
	repo repository.TaskList
}

// NewTaskListService initializes a new TaskListService with the provided repository.
func NewTaskListService(repo repository.TaskList) *TaskListService {
	return &TaskListService{
		repo: repo,
	}
}

// Create creates a new task list for the specified user.
func (s *TaskListService) Create(userId string, list model.TaskList) (int, error) {
	if err := validateCreateTaskList(list); err != nil {
		return 0, err
	}
	return s.repo.Create(userId, list)
}

// GetAll retrieves all task lists for the specified user.
func (s *TaskListService) GetAll(userId string) ([]model.TaskList, error) {
	return s.repo.GetAll(userId)
}

// GetById retrieves a specific task list by its ID for the specified user.
func (s *TaskListService) GetById(userId string, listId int) (model.TaskList, error) {
	return s.repo.GetById(userId, listId)
}

// Delete deletes a specific task list by its ID for the specified user.
func (s *TaskListService) Delete(userId string, listId int) error {
	if err := validateDeleteTaskList(listId); err != nil {
		return err
	}
	return s.repo.Delete(userId, listId)
}

// Update updates a specific task list by its ID for the specified user.
func (s *TaskListService) Update(userId string, listId int, input model.UpdateTaskListInput) error {
	if err := validateUpdateTaskList(input); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

// validateCreateTaskList checks if the task list is valid for creation.
func validateCreateTaskList(list model.TaskList) error {
	if list.Title == "" {
		return errors.New("task list name cannot be empty")
	}
	return nil
}

// validateUpdateTaskList checks if the update input is valid.
func validateUpdateTaskList(input model.UpdateTaskListInput) error {
	if input.Title == nil && input.Description == nil {
		return errors.New("no fields to update")
	}
	return nil
}

// validateDeleteTaskList checks if the list ID is valid for deletion.
func validateDeleteTaskList(listId int) error {
	if listId <= 0 {
		return errors.New("invalid task list ID")
	}
	return nil
}
