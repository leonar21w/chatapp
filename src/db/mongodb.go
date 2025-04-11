package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func GetMongoClient() (*mongo.Client, error) {
	var err error

	clientOnce.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatal("MONGODB_URI environment variable is not set")
		}
		fmt.Print("Connecting to MongoDB...")
		fmt.Println(uri)

		clientOptions := options.Client().ApplyURI(uri)

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// ✅ Correct argument order
		clientInstance, err = mongo.Connect(clientOptions)
		if err != nil {
			log.Fatalf("Failed to create MongoDB client: %v", err)
		}

		if err = clientInstance.Ping(ctx, nil); err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		log.Println("✅ Successfully connected to MongoDB")
	})

	return clientInstance, err
}
