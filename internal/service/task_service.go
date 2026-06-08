package service

import (
	"errors"
	"github.com/temurova-ui/todo_api/internal/models"
	"github.com/temurova-ui/todo_api/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService{
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(task models.Task)(models.Task, error){
	if task.Title == ""{
		return task, errors.New("title is required")
	}
	return s.repo.Create(task)
}
	


