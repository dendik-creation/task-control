package http

import (
	"log"
	"net/http"

	"github.com/dendik-creation/task-control/internal/domain"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase domain.UserUsecase
}

func NewAuthHandler(r *gin.Engine, us domain.UserUsecase) {
	handler := &AuthHandler{
		authUsecase: us,
	}

	api := r.Group("/api")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format input"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	err := h.authUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		log.Printf("Error registering user: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	log.Printf("Received body request: %v", c.Request.Body)
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	token, err := h.authUsecase.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
