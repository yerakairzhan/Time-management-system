package handler

import (
	"TimeManagementSystem/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", h.createTask)
	router.GET("/tasks", h.listTasks)
}

func (h *TaskHandler) createTask(c *gin.Context) {
	// логика создания задачи
}

func (h *TaskHandler) listTasks(c *gin.Context) {
	// логика получения списка задач
}
