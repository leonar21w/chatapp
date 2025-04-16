package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connection struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User_A      primitive.ObjectID `bson:"user_a" json:"user_a"`
	User_B      primitive.ObjectID `bson:"user_b" json:"user_b"`
	Status      int                `bson:"status" json:"status"` //
	RequestedAt time.Time          `bson:"requested_at" json:"requested_at"`
	ResponseAt  time.Time          `bson:"response_at" json:"response_at"`
}

type FriendRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FromID    primitive.ObjectID `bson:"from_id" json:"from_id"`
	TargetID  primitive.ObjectID `bson:"target_id" json:"target_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	Status    int                `bson:"status" json:"status"` // 0: pending, 1: accepted, 2: rejected
}

type FriendRequestForClient struct {
	User      UserNotSensitive
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	Status    int       `bson:"status" json:"status"` // 0: pending, 1: accepted, 2: rejected
}
