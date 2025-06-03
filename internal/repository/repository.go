package repository

import (
	"TaskManager/internal/domain/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Authorization defines the interface for user authentication and authorization operations.
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

// TaskList defines the interface for task list operations.
type TaskList interface {
	Create(userId string, list model.TaskList) (int, error)
	GetAll(userId string) ([]model.TaskList, error)
	GetById(userId string, listId int) (model.TaskList, error)
	Delete(userId string, listId int) error
	Update(userId string, listId int, input model.UpdateTaskListInput) error
}

// Repository defines the interface for interacting with the data layer.
type Repository struct {
	Authorization
	TaskList
}

// NewRepository initializes a new Repository instance with MongoDB implementations.
func NewRepository(client *mongo.Client, dbName string) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(client, dbName),
		TaskList:      NewTaskListMongo(client, dbName),
	}
}
