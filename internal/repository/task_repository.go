package repository

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/temurova-ui/todo_api/internal/models"
)

type TaskRepository interface {
    Create(task models.Task)(models.Task, error)
    GetAll()([]models.Task, error)
    GetByID(id int)(models.Task, error)
    Update(id int, task models.Task)(models.Task, error)
    Delete(id int) error
}

type taskRepository struct {
    filePath string
}

func NewTaskRepository(path string) TaskRepository{
	return &taskRepository{filePath: path,}
}

func (r *taskRepository) loadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) saveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *taskRepository) Create(task models.Task) (models.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return models.Task{}, err
	}

	maxID := 0

	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	task.ID = maxID + 1
	
	task.CreatedAt = time.Now()
	
	tasks = append(tasks, task)

	err = r.saveTasks(tasks)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) GetAll() ([]models.Task, error) {
	return r.loadTasks()
}

func (r *taskRepository) GetByID(id int) (models.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return models.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("task not found")
}

func (r *taskRepository)Update(id int, updatedTask models.Task)(models.Task, error){
	tasks, err := r.loadTasks()
	if err != nil{
		return models.Task{}, err
	}
	for i, task := range tasks{
		if task.ID == id{
			updatedTask.ID = id
			updatedTask.CreatedAt = task.CreatedAt

			tasks[i] = updatedTask

			err = r.saveTasks(tasks)
			if err != nil{
				return models.Task{}, err
			}
			return updatedTask, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (r *taskRepository)Delete(id int)error{
	tasks, err := r.loadTasks()
	if err != nil{
		return err
	}

	for i, task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			
			return r.saveTasks(tasks)
		}
	}
	return errors.New("task not found")
}