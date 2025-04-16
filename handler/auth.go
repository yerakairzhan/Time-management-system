package handler

import (
	db "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type signUpInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var user db.User
	user.Email = sql.NullString{String: strings.ToLower(input.Email), Valid: true}
	user.HashedPassword = sql.NullString{String: service.GeneratePasswordHash(input.Password), Valid: true}

	id, err := h.authService.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) logIn(c *gin.Context) {
}
