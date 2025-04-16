package entry

import (
	"context"

	"github.com/leonar21w/chat-backend/src/db/repository"
	friendRequests "github.com/leonar21w/chat-backend/src/db/repository/friend_requests"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
	UserRepo          *repository.UserRepo
	FriendRequestRepo *friendRequests.FriendRequestRepo
	ConnectionRepo    *repository.ConnectionRepo
}

func RepoInit(ctx context.Context, client *mongo.Client) (*Repositories, error) {

	userRepo := repository.NewUserRepo(client)
	friendRequestRepo := friendRequests.NewFriendRequestRepo(client)
	connectionRepo := repository.NewConnectionRepo(client)

	if err := userRepo.EnsureUserIndexes(ctx); err != nil {
		return nil, err
	}
	if err := friendRequestRepo.EnsureFriendRequestIndexes(ctx); err != nil {
		return nil, err
	}

	return &Repositories{
		UserRepo:          userRepo,
		FriendRequestRepo: friendRequestRepo,
		ConnectionRepo:    connectionRepo,
	}, nil

}
