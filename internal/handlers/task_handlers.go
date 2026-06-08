package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/temurova-ui/todo_api/internal/models"
	"github.com/temurova-ui/todo_api/internal/service"
)

type TaskHandler struct{
    service *service.TaskService
}

func NewTaskHandler(service *service.TaskService)*TaskHandler{
    return &TaskHandler{
        service: service,
    }
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter,r *http.Request){
    var task models.Task

    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    createdTask, err := h.service.CreateTask(task)
    if err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.Header().Set(
        "Content-Type",
        "application/json",
    )
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdTask)


}
