package repository

import "TaskManager/internal/domain/model"

// Authorization defines the interface for user authentication and authorization operations.
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tokenString string) (int, error)
}

// TaskList defines the interface for task list operations.
type TaskList interface {
	Create(userId int, list model.TaskList) (int, error)
	GetAll(userId int) ([]model.TaskList, error)
	GetById(userId int, listId int) (model.TaskList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input model.UpdateTaskListInput) error
}

// Repository defines the interface for interacting with the data layer.
type Repository struct {
	Authorization
	TaskList
}

// NewRepository initializes a new Repository instance with MongoDB implementations.
func NewRepository() *Repository {
	return &Repository{
		Authorization: NewAuthMongo(),
		TaskList:      NewTaskListMongo(),
	}
}
