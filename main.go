package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"gin-gorm-crud/config"
	"gin-gorm-crud/controller"
	"gin-gorm-crud/utils"
	"gin-gorm-crud/repository"
	"gin-gorm-crud/router"
	"gin-gorm-crud/service"
	
	"time"
	"log"
)

func main() {
	//Database
	db:= config.DatabaseConnection()
	validate := validator.New()

	//Init Repository
	taskRepository := repository.NewTaskRepositoryImpl(db)

	//Init Service
	taskService, err := service.NewTaskServiceImpl(taskRepository, validate)
	if err != nil {
		// Handle error appropriately, such as logging and exiting
		log.Fatalf("Error initializing task service: %v", err)
	}

	//Init controller
	taskController := controller.NewTaskController(taskService)

	//Router
	routes := router.TaskRouter(taskController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)

}