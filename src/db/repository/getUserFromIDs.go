package repository

import (
	"context"
	"fmt"
	model "github.com/leonar21w/chat-backend/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

func GetUserUnsensitiveFromIDs(ctx context.Context, userRepo *UserRepo, ids []primitive.ObjectID) ([]model.UserNotSensitive, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := userRepo.col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding users: %w", err)
	}
	defer cursor.Close(ctx)

	var result []model.UserNotSensitive

	for cursor.Next(ctx) {
		var u model.User
		if err := cursor.Decode(&u); err != nil {
			log.Println("error decoding user:", err)
			continue
		}

		result = append(result, sanitizeUser(u))
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return result, nil
}

func sanitizeUser(u model.User) model.UserNotSensitive {
	return model.UserNotSensitive{
		ID:              u.ID,
		Username:        u.Username,
		Name:            u.Name,
		ProfileImageURL: u.ProfileImageURL,
	}
}
