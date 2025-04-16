package main

import (
	"log"
	"net/http"

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

	authService := service.NewAuthService(userRepo)
	taskService := service.NewTaskService(taskRepo, userRepo)

	httpHandler := handler.NewHandler(taskService, authService)

	// Сервер
	srv := &http.Server{
		Addr:    ":8080",
		Handler: httpHandler.InitRoutes(), // правильно инициализируем маршруты
	}

	log.Println("🚀 Server is running on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
