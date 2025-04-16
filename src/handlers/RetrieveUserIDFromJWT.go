package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserIDFromJWT(c *gin.Context) (primitive.ObjectID, error) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		return primitive.NilObjectID, errors.New("userID doesnt exist in request")
	}

	userIDStr, ok := userIDRaw.(string)
	if !ok {
		return primitive.NilObjectID, errors.New("did not receive userID")
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return primitive.NilObjectID, errors.New("could not read userID, bad format")
	}

	return userID, nil
}
