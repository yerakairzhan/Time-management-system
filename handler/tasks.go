package handler

import (
	db "TimeManagementSystem/db/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TaskInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Priority    string `json:"priority"` // или string, если low/medium/high
	Deadline    string `json:"deadline"`
}

type UpdateTaskInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
	Priority    *string `json:"priority"`
	Deadline    *string `json:"deadline"`
}

func (h *Handler) createTask(c *gin.Context) {
	var input TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректные данные задачи"})
		return
	}

	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	deadline, err := time.Parse(time.RFC3339, input.Deadline)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверный формат дедлайна (ожидается RFC3339)")
		return
	}

	task := db.Task{
		UserID:      int32(userId),
		Name:        input.Name,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Category:    input.Category,
		Priority:    input.Priority,
		Deadline:    sql.NullTime{Time: deadline, Valid: true},
	}

	var id int
	if id, err = h.taskService.CreateTask(userId, task); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "не удалось создать задачу")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTasks(c *gin.Context) {
	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	tasksOfUser, err := h.taskService.GetTasksByUserID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasksOfUser)
}

func (h *Handler) getTask(c *gin.Context) {
	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.taskService.GetTaskById(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if task.UserID != int32(userId) {
		newErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) updateTask(c *gin.Context) {
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.taskService.GetTaskById(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if task.UserID != int32(userId) {
		newErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}

	if input.Name != nil {
		task.Name = *input.Name
	}
	if input.Description != nil {
		task.Description = sql.NullString{String: *input.Description, Valid: true}
	}
	if input.Category != nil {
		task.Category = *input.Category
	}
	if input.Priority != nil {
		task.Priority = *input.Priority
	}
	if input.Deadline != nil {
		parsedDeadline, err := time.Parse(time.RFC3339, *input.Deadline)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deadline format"})
			return
		}
		task.Deadline = sql.NullTime{Time: parsedDeadline, Valid: true}
	}

	err = h.taskService.UpdateTask(taskId, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.taskService.GetTaskById(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if task.UserID != int32(userId) {
		newErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}

	err = h.taskService.DeleteTask(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getUserIdFromContext(c *gin.Context) (int, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return 0, errors.New("пустой заголовок авторизации")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("неверный формат авторизации")
	}

	return h.authService.ParseToken(parts[1])
}
