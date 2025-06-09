package api

import (
	"log"
	"net/http"

	"github.com/flames31/jobqueue/internal/auth"
	"github.com/flames31/jobqueue/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) POSTUserRegister(c *gin.Context) {
	var req struct {
		Email    string
		Password string
	}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("Error registering user : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error registering user : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.Service.UserService.CreateUser(&model.User{
		Email:        req.Email,
		PasswordHash: hash,
	})
	if err != nil {
		log.Printf("Error registering user : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

func (h *handler) POSTUserLogin(c *gin.Context) {
	var req struct {
		Email    string
		Password string
	}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("Error logging in : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := h.Service.UserService.VerifyCredentials(req.Email, req.Password)
	if err != nil {
		log.Printf("Error registering user : %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect password",
		})
		return
	}

	token, err := auth.NewJWT(userID, h.JWTSecret)
	if err != nil {
		log.Printf("Error registering user : %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    userID,
		"email": req.Email,
		"token": token,
	})
}
