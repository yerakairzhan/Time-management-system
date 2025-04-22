package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) startTaskTimer(c *gin.Context) {
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

	err = h.taskService.StartTask(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task timer started successfully"})
}

func (h *Handler) stopTaskTimer(c *gin.Context) {
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

	err = h.taskService.StopTask(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task timer stopped successfully"})
}

func (h *Handler) getTaskTime(c *gin.Context) {
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

	time, err := h.taskService.GetTimeSpent(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var formatted []string
	for _, log := range time {
		formatted = append(formatted, fmt.Sprintf("%02dh %02dm %02ds", int(log.Hours()), int(log.Minutes())%60, int(log.Seconds())%60))
	}

	c.JSON(http.StatusOK, gin.H{"message": formatted})
}

func (h *Handler) getTaskTimeHistory(c *gin.Context) {
}
