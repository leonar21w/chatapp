package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`      //Accesing the message in mongo
	Text      string             `bson:"text" json:"text"`             //text content
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"` //show time

	SenderID primitive.ObjectID `bson:"sender_id" json:"sender_id"` //show who sent the message this will be user id
	GroupID  primitive.ObjectID `bson:"group_id" json:"group_id"`   //show what conversation it is a part of
}
