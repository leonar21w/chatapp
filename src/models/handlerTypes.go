package model

type RegisterRequest struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ProfileImageURL string `json:"profile_image_url"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
