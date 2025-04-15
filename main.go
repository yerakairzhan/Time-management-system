package main

import (
	"TimeManagementSystem/config"
	db "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/handler"
	"TimeManagementSystem/repository/postgres"
	"TimeManagementSystem/service"
	"log"
	"net/http"
)

func main() {
	dbConn := config.DatabaseConnection()
	defer dbConn.Close()

	queries := db.New(dbConn)

	userRepo := postgres.NewUserRepository(queries)
	taskRepo := postgres.NewTaskRepository(queries)

	taskService := service.NewTaskService(taskRepo, userRepo)

	httpHandler := handler.NewHTTPHandler(taskService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: Initroutes.,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
