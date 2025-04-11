package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	model "github.com/leonar21w/chat-backend/src/models"
	"github.com/leonar21w/chat-backend/src/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LoginRequest(user *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.LoginRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		user, err := user.FindByEmail(c, input.Email)

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding email"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}

		//give clients token -> clients will store this token in their userdefaults if they lose it they can just login again

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": user.ID.Hex(),
			"exp":    time.Now().Add(72 * time.Hour).Unix(),
		})

		secret := os.Getenv("JWT_SECRET")
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

//For now we can just ignore the token until it self deletes, if we want to logoff on the client we can delete it there
