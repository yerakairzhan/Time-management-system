package handler

import (
	"TimeManagementSystem/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskService         service.TaskService
	authService         service.Authorization
	notificationService service.NotificationService
}

func NewHandler(taskService service.TaskService, authService service.Authorization, notificationService service.NotificationService) *Handler {
	return &Handler{
		taskService:         taskService,
		authService:         authService,
		notificationService: notificationService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/log-in", h.logIn)
	}

	api := router.Group("/api")
	{
		api.GET("/tasks", h.getTasks)          // Get all tasks
		api.POST("/tasks", h.createTask)       // Create a new task
		api.GET("/tasks/:id", h.getTask)       // Get a specific task by ID
		api.PUT("/tasks/:id", h.updateTask)    // Update a specific task by ID
		api.DELETE("/tasks/:id", h.deleteTask) // Delete a specific task by ID

		api.POST("/tasks/:id/timer/start", h.startTaskTimer) // Start timer for a task
		api.POST("/tasks/:id/timer/stop", h.stopTaskTimer)   // Stop timer for a task
		api.GET("/tasks/:id/time", h.getTaskTime)            // Get the time spent on a task

		api.POST("/notifications", h.createNotification)       // Create a new notification
		api.GET("/notifications", h.getNotifications)          // Get all notifications
		api.PUT("/notifications/:id", h.updateNotification)    // Update a specific notification
		api.DELETE("/notifications/:id", h.deleteNotification) // Delete a specific notification

		api.GET("/reports/time-spent", h.getTimeSpentReport)                          // Get a report on time spent on tasks
		api.GET("/reports/task-completion", h.getTaskCompletionReport)                // Get a report on task completion rates
		api.GET("/reports/completion-on-time", h.getCompletionOnTimeReport)           // Get a report on tasks completed on time
		api.GET("/reports/average-completion-time", h.getAverageCompletionTimeReport) // Get a report on average task completion time
	}

	return router
}
