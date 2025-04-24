package handler

import (
	db "TimeManagementSystem/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NotificationInput struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (h *Handler) createNotification(c *gin.Context) {
	var input NotificationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректные данные задачи"})
		return
	}

	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notification := db.Notification{
		UserID:  sql.NullInt32{Int32: int32(userId), Valid: true},
		Title:   sql.NullString{String: input.Title, Valid: true},
		Message: sql.NullString{String: input.Message, Valid: true},
	}

	var id int
	if id, err = h.notificationService.CreateNotification(notification); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "не удалось создать Напоминание")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getNotifications(c *gin.Context) {
	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notifications, err := h.notificationService.GetNotificationsByUserID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func (h *Handler) updateNotification(c *gin.Context) {
	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notificationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid notification id")
		return
	}

	notification, err := h.notificationService.GetById(notificationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if notification.UserID.Int32 != int32(userId) {
		newErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}

	var input NotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification.Title = sql.NullString{String: input.Title, Valid: true}
	notification.Message = sql.NullString{String: input.Message, Valid: true}

	err = h.notificationService.UpdateNotification(notification)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification updated"})
}

func (h *Handler) deleteNotification(c *gin.Context) {
	userId, err := h.getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notificationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid notification id")
		return
	}

	notification, err := h.notificationService.GetById(notificationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if notification.UserID.Int32 != int32(userId) {
		newErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}

	err = h.notificationService.DeleteNotification(notificationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}
