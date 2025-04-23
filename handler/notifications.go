package handler

import (
	db "TimeManagementSystem/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
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
}

func (h *Handler) updateNotification(c *gin.Context) {
}

func (h *Handler) deleteNotification(c *gin.Context) {
}
