package main

import (
	"database/sql"
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
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)

	queries := db.New(dbConn)

	userRepo := postgres.NewUserRepository(queries)
	taskRepo := postgres.NewTaskRepository(queries)

	authService := service.NewAuthService(userRepo)
	taskService := service.NewTaskService(taskRepo, userRepo)

	httpHandler := handler.NewHandler(taskService, authService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: httpHandler.InitRoutes(),
	}

	log.Println("🚀 Server is running on http://localhost" + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
