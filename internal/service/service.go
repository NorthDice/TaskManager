package service

import (
	"TaskManager/internal/domain/model"
	"TaskManager/internal/repository"
)

// Authorization defines the interface for user authentication and authorization operations.
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tokenString string) (string, error)
}

// TaskList defines the interface for task list operations.
type TaskList interface {
	Create(userId int, list model.TaskList) (int, error)
	GetAll(userId int) ([]model.TaskList, error)
	GetById(userId int, listId int) (model.TaskList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input model.UpdateTaskListInput) error
}

// Service defines the interface for the service layer, combining authorization and task list operations.
type Service struct {
	Authorization
	TaskList
}

// NewService initializes a new Service instance with the provided repository.
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TaskList:      NewTaskListService(repo.TaskList),
	}
}
