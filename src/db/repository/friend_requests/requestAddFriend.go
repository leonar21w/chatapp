package friendRequests

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type FriendRequestRepo struct {
	col *mongo.Collection
}

func NewFriendRequestRepo(client *mongo.Client) *FriendRequestRepo {
	return &FriendRequestRepo{
		col: client.Database("chatapp").Collection("friend_requests"),
	}
}

func (r *FriendRequestRepo) RequestAddFriend(ctx context.Context, target primitive.ObjectID, initiator primitive.ObjectID) error {
	if target == initiator {
		return errors.New("you cannot send a friend request to yourself")
	}

	filter := bson.M{
		"from_id":   initiator,
		"target_id": target,
		"status":    0,
	}

	update := bson.M{
		"$set": bson.M{
			"created_at": time.Now(),
		},
		"$setOnInsert": bson.M{
			"_id": primitive.NewObjectID(),
		},
	}

	opts := options.UpdateOne().SetUpsert(true)

	_, err := r.col.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *FriendRequestRepo) EnsureFriendRequestIndexes(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "from_id", Value: 1},
			{Key: "target_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.col.Indexes().CreateOne(ctx, indexModel)
	return err
}
