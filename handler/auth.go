package handler

import (
	db "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input SignInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var user db.User
	user.Email = input.Email
	user.HashedPassword = service.GeneratePasswordHash(input.Password)

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
	var input SignInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.authService.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
