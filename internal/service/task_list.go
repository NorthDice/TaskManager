package service

import "TaskManager/internal/repository"

type TaskListService struct {
	repo repository.TaskList
}

func NewTaskListService(repo repository.TaskList) *TaskListService {
	return &TaskListService{
		repo: repo,
	}
}
