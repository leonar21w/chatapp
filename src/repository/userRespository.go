package repository

import (
	"context"

	model "github.com/leonar21w/chat-backend/src/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepo struct {
	col *mongo.Collection
}

func NewUserRepo(client *mongo.Client) *UserRepo {
	return &UserRepo{
		col: client.Database("chatapp").Collection("users"),
	}
}

func (r *UserRepo) EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true).SetName("unique_email"),
		},
		{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true).SetName("unique_username"),
		},
	}
	_, err := r.col.Indexes().CreateMany(ctx, indexes)
	return err
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	filter := bson.M{"email": email}

	var user model.User
	err := r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil //No user found under this email
		}
		//An error happened
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindByHandle(ctx context.Context, username string) (*model.User, error) {
	filter := bson.M{"username": username}

	var user model.User
	err := r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No emails found
		}
		return nil, err // error occured
	}
	return &user, nil
}
