package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonar21w/chat-backend/src/db/repository"
	friendRequests "github.com/leonar21w/chat-backend/src/db/repository/friend_requests"
	model "github.com/leonar21w/chat-backend/src/models"
)

func SendFriendRequestHandler(requestRepo *friendRequests.FriendRequestRepo, userRepo *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.AddFriendRequest

		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request" + err.Error()})
			return
		}

		receiver, err := userRepo.FindByHandle(c, request.Receiver)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error finding user" + err.Error()})
			return
		}

		if receiver == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This user does not exist, check username query"})
			return
		}

		sender, err := userRepo.FindByHandle(c, request.Sender)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error finding user data" + err.Error()})
			return
		}

		if sender == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Something is wrong with the sender data, make sure you are logged in"})
			return
		}

		err = requestRepo.RequestAddFriend(c, receiver.ID, sender.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error sending friend request" + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Friend request sent successfully"})
	}
}
