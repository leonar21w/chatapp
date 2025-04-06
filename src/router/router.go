package router

import (
	"github.com/gin-gonic/gin"
	auth "github.com/leonar21w/chat-backend/src/handlers/Auth"
	"github.com/leonar21w/chat-backend/src/middleware"
	"github.com/leonar21w/chat-backend/src/repository"
)

func Setup(r *gin.Engine, repo *repository.UserRepository) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api/v1/lake")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/user/auth/register", auth.CreateUser(repo))
	}

}
