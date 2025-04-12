package model

type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ProfileImageURL string `json:"profile_image_url"` // optional
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AddFriendRequest struct {
	Receiver string `json:"receiver" binding:"required"`
	Sender   string `json:"sender" binding:"required"`
}
