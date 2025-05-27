package service

import (
	"TaskManager/internal/domain/model"
	"TaskManager/internal/repository"
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
func (s *TaskListService) Create(userId int, list model.TaskList) (int, error) {
	return s.repo.Create(userId, list)
}

// GetAll retrieves all task lists for the specified user.
func (s *TaskListService) GetAll(userId int) ([]model.TaskList, error) {
	return s.repo.GetAll(userId)
}

// GetById retrieves a specific task list by its ID for the specified user.
func (s *TaskListService) GetById(userId int, listId int) (model.TaskList, error) {
	return s.repo.GetById(userId, listId)
}

// Delete deletes a specific task list by its ID for the specified user.
func (s *TaskListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

// Update updates a specific task list by its ID for the specified user.
func (s *TaskListService) Update(userId int, listId int, input model.UpdateTaskListInput) error {
	return s.repo.Update(userId, listId, input)
}
