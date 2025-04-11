package repository

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ConnectionRepo struct {
	col *mongo.Collection
}

func NewConnectionRepo(client *mongo.Client) *ConnectionRepo {
	return &ConnectionRepo{
		col: client.Database("chatapp").Collection("connections"),
	}
}

//Add friend, search Friend, remove friend, get friends, block friend, unblock friend

//friend requests lookup, accept, deny, block, unblock
