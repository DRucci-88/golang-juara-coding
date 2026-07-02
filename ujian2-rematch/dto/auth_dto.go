package dto

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `jsong:"password" binding:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}