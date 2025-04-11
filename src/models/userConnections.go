package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connection struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User_A      primitive.ObjectID `bson:"user_a" json:"user_a"`
	User_B      primitive.ObjectID `bson:"user_b" json:"user_b"`
	Status      int                `bson:"status" json:"status"` //-1 rejected, 0 pending, 1 accepted
	RequestedAt time.Time          `bson:"requested_at" json:"requested_at"`
	ResponseAt  time.Time          `bson:"response_at" json:"response_at"`
}
