package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonar21w/chat-backend/src/db/repository"
	"github.com/leonar21w/chat-backend/src/handlers"
	"github.com/leonar21w/chat-backend/src/middleware"
)

func Setup(r *gin.Engine, repo *repository.UserRepo) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api/v1/auth")
	api.POST("/register", handlers.RegisterNewUser(repo))
	api.POST("/login", handlers.LoginRequest(repo))

	apiAuthenticated := r.Group("/api/v1")
	apiAuthenticated.Use(middleware.AuthMiddleware()) //
	{
	}

}
