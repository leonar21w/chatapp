package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//A group is a conversation between multiple users, a 1-1 chat is also considered a group

type Group struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`                //Accesing the conversation in mongo
	Name           string               `bson:"name" json:"name"`                       //Name of the group
	Users          []primitive.ObjectID `bson:"users" json:"users"`                     //Users in the group
	GroupFunctions bool                 `bson:"group_functions" json:"group_functions"` //if true then its a group chat, else its a 1-1 chat
	GroupImageURL  string               `bson:"group_image_url" json:"group_image_url"` //group image url
}
