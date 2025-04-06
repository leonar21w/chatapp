package repository

import (
	"context"

	model "github.com/leonar21w/chat-backend/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{col: client.Database("chat").Collection("users")}
}

func (r *UserRepository) CreateUser(user *model.User, ctx context.Context) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	_, err = r.col.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindUserByEmail(email string, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindUserByID(id primitive.ObjectID, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
