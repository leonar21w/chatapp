package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonar21w/chat-backend/src/db/repository"
	model "github.com/leonar21w/chat-backend/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNewUser(userRepo *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.RegisterRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		exists, _ := userRepo.FindByEmail(c, request.Email)
		if exists != nil {
			c.JSON(400, gin.H{"error": "User with this email already exists"})
			return
		}

		exists, _ = userRepo.FindByHandle(c, request.Username)
		if exists != nil {
			c.JSON(400, gin.H{"error": "User with this username already exists"})
			return
		}

		hashP, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

		user := &model.User{
			ID:       primitive.NewObjectID(),
			Username: request.Username,
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashP),
		}

		if err := userRepo.InsertNewUser(c, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}
