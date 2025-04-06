package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Email           string             `bson:"email" json:"email"`
	ProfileImageURL string             `bson:"profile_image_url" json:"profile_image_url"`
	Password        string             `bson:"password" json:"password"`
}

//Create new user -> Post/users
