package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonar21w/chat-backend/main/entry"
	"github.com/leonar21w/chat-backend/src/handlers"
	"github.com/leonar21w/chat-backend/src/middleware"
)

func Setup(r *gin.Engine, repositories *entry.Repositories) {

	fqRepo := repositories.FriendRequestRepo
	userRepo := repositories.UserRepo
	//connection := repositories.ConnectionRepo

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	api := r.Group("/api/v1/auth")
	api.POST("/register", handlers.RegisterNewUser(repositories.UserRepo))
	api.POST("/login", handlers.LoginRequest(repositories.UserRepo))

	apiAuthenticated := r.Group("/api/v1")
	apiAuthenticated.Use(middleware.AuthMiddleware()) //
	{
		apiAuthenticated.POST("/friend_request/create", handlers.SendFriendRequestHandler(fqRepo, userRepo))
		apiAuthenticated.GET("/friend_request/fqList", handlers.GetAllFriendRequestsHandler(fqRepo, userRepo))
	}

}
