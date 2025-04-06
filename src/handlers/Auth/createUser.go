package auth

//creates new user in the db from client request

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/leonar21w/chat-backend/src/models"
	"github.com/leonar21w/chat-backend/src/repository"
)

func CreateUser(repository *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newUser model.User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad client input" + err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		existingUser, _ := repository.FindUserByEmail(newUser.Email, ctx)
		if existingUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already in use, use a different email"})
			return
		}

		existingHandle, _ := repository.FindUserByID(newUser.ID, ctx)
		if existingHandle != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Handle already in use, use a different handle"})
			return
		}

		err := repository.CreateUser(&newUser, ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user" + err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User created in db"})
	}
}
