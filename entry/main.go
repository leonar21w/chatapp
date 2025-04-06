package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leonar21w/chat-backend/src/db"
	"github.com/leonar21w/chat-backend/src/repository"
	"github.com/leonar21w/chat-backend/src/router"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)

	//Database

	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Couldn't disconnect from database: %v", err)
		}
	}()

	//server
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	router.Setup(r, repository.NewUserRepository(client))

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Couldn't run server: %v", err)
	}

}
