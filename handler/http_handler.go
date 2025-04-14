package handler

import (
	"TimeManagementSystem/service"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service *service.TaskService
}

func NewHTTPHandler(s *service.TaskService) *HTTPHandler {
	return &HTTPHandler{service: s}
}

func (h *HTTPHandler) Start() {
	r := gin.Default()

	r.POST("/tasks", h.createTask)
	r.GET("/tasks", h.listTasks)

	r.Run(":8080")
}
