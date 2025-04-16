package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonar21w/chat-backend/src/db/repository"
	friendRequests "github.com/leonar21w/chat-backend/src/db/repository/friend_requests"
)

// Handler for getting a list of friend requests, the client will ask for this data. expected { list of users that send our USERID a friendrequest }
func GetAllFriendRequestsHandler(fqRepo *friendRequests.FriendRequestRepo, userRepo *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := GetUserIDFromJWT(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error processing userID" + err.Error()})
			return
		}

		//Get all friend request statuses
		friendRequests, err := fqRepo.GetFriendRequestsForClient(c, userID, userRepo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting all friend requests" + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"friendRequests": friendRequests})
	}
}
