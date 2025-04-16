package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserContextID(c *gin.Context) (primitive.ObjectID, error) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		return primitive.NilObjectID, errors.New("user ID not found in context")
	}

	userID, ok := userIDRaw.(primitive.ObjectID)
	if !ok || userID == primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("invalid user ID")
	}

	return userID, nil
}
