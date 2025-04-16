package friendRequests

import (
	"context"
	"github.com/leonar21w/chat-backend/src/db/repository"
	model "github.com/leonar21w/chat-backend/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

// Returns a friendrequest list for the userID that the client sends
func (r *FriendRequestRepo) GetFriendRequestsForClient(ctx context.Context, userID primitive.ObjectID, userRepo *repository.UserRepo) ([]model.FriendRequestForClient, error) {
	filter := bson.M{"target_id": userID, "status": 0} //0 is pending, 1 is accepted, if rejected delete the data

	pendingFriendRequests, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer pendingFriendRequests.Close(ctx)
	//sender id, -> search user safe -> return with statuses

	var requests []model.FriendRequest
	var requesterIDs []primitive.ObjectID

	for pendingFriendRequests.Next(ctx) {
		var temp model.FriendRequest
		if err := pendingFriendRequests.Decode(&temp); err != nil {
			log.Println(err)
			continue
		}
		requests = append(requests, temp)
		requesterIDs = append(requesterIDs, temp.FromID)
	}

	users, err := repository.GetUserUnsensitiveFromIDs(ctx, userRepo, requesterIDs)
	if err != nil {
		return nil, err
	}

	mapUsers := make(map[primitive.ObjectID]model.UserNotSensitive)
	for _, u := range users {
		mapUsers[u.ID] = u
	}

	//merge
	var res []model.FriendRequestForClient
	for _, request := range requests {
		user, ok := mapUsers[request.FromID]
		if !ok {
			log.Println("error with friend request merge")
			continue
		}
		res = append(res, model.FriendRequestForClient{
			User:      user,
			CreatedAt: request.CreatedAt,
			Status:    request.Status,
		})
	}
	return res, nil
}
