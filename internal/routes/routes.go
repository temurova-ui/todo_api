package routes

import(
	"net/http"
	"github.com/temurova-ui/todo_api/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux, handler *handlers.TaskHandler){
	mux.HandleFunc("POST/tasks", handler.CreateTask)
}
