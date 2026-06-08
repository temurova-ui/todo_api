package main

import (
	"log"
	"net/http"

	"github.com/temurova-ui/todo_api/internal/handlers"
	"github.com/temurova-ui/todo_api/internal/middleware"
	"github.com/temurova-ui/todo_api/internal/repository"
	"github.com/temurova-ui/todo_api/internal/routes"
	"github.com/temurova-ui/todo_api/internal/service"
)

func main(){
repo := repository.NewTaskRepository("data/tasks.json")

service := service.NewTaskService(repo)

handler := handlers.NewTaskHandler(service)

mux := http.NewServeMux()

routes.RegisterRoutes(mux, handler)

logged := middleware.Logger(mux)

log.Println("Server started on :8080")

err := http.ListenAndServe(":8080", logged)
	if err != nil{
		log.Fatal(err)
	}
}