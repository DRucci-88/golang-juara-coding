package dto

import "ujian2_rematch/model"

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func NewUserResponse(model *model.User) *UserResponse {
	if model == nil {
		return nil
	}

	return &UserResponse{
		ID:    model.ID,
		Email: model.Email,
	}
}
