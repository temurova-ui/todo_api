package main

import (
	"log"
	"net/http"

	"github.com/temurova-ui/todo_api/internal/config"
	"github.com/temurova-ui/todo_api/internal/handlers"
	"github.com/temurova-ui/todo_api/internal/middleware"
	"github.com/temurova-ui/todo_api/internal/repository"
	"github.com/temurova-ui/todo_api/internal/routes"
	"github.com/temurova-ui/todo_api/internal/service"
)

func main() {
	cfg, err := config.New("./internal/config/config.env")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTaskRepository(cfg.Storage)

	taskService := service.NewTaskService(repo)

	taskHandler := handlers.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, taskHandler)

	logged := middleware.Logger(mux)

	log.Println("Server started on", cfg.HttpPort)

	err = http.ListenAndServe(cfg.HttpPort, logged)
	if err != nil {
		log.Fatal(err)
	}
}