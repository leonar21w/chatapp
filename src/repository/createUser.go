package repository

import (
	"context"

	model "github.com/leonar21w/chat-backend/src/models"
)

func (r *UserRepo) InsertNewUser(ctx context.Context, user *model.User) error {
	_, err := r.col.InsertOne(ctx, user)
	return err
}
