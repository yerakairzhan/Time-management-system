package main

import (
	"TimeManagementSystem/config"
	db "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/handler"
	"TimeManagementSystem/repository/postgres"
	"TimeManagementSystem/service"
)

func main() {
	dbConn := config.DatabaseConnection()
	defer dbConn.Close()

	queries := db.New(dbConn)

	userRepo := postgres.NewUserRepository(queries)
	taskRepo := postgres.NewTaskRepository(queries)

	taskService := service.NewTaskService(taskRepo, userRepo)

	httpHandler := handler.NewHTTPHandler(taskService)

	httpHandler.Start()

}
