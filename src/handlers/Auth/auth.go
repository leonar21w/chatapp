package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"github.com/leonar21w/chat-backend/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	UserID   string `json:"user_id" binding:"required_without=Email"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"required"`
}

var jwtSecret = []byte("JWTSECRET")

func LoginHandler(repo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(400, gin.H{"error": "Bad client input" + err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user, err := repo.FindUserByEmail(loginRequest.Email, ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with DB" + err.Error()})
			return
		}

		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": user.ID,
			"exp":    time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token" + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
