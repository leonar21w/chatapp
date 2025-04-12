package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`            //Accesing the user in mongo
	Username        string             `bson:"username" json:"username"` //user handle like @Leonard, will be unique for all users
	Name            string             `bson:"name" json:"name"`         //username can conflict
	Email           string             `bson:"email" json:"email"`       //user emails must be unique for all users
	ProfileImageURL string             `bson:"profile_image_url" json:"profile_image_url"`
	Password        string             `bson:"password" json:"password"` //Hashed password
	LastOnline      time.Time          `bson:"last_online" json:"last_online"`
	Status          string             `bson:"status" json:"status"` //user status, (bio)

}

//The user struct that we should submit to the client will be implemented later, data like last online needs to update
//Sensitive data like password and ID should not be sent to the client
