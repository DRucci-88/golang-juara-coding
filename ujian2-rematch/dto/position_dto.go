package dto

import "ujian2_rematch/model"

type PositionCreateRequest struct {
	Title      string  `json:"title" binding:"required,max=50"`
	BaseSalary float64 `json:"base_salary" binding:"required"`
}

type PositionUpdateRequest struct {
	Title      string  `json:"title" binding:"required,max=50"`
	BaseSalary float64 `json:"base_salary" binding:"required"`
}

type PositionResponse struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	BaseSalary float64 `json:"base_salary"`
}

func NewPositionResponse(model *model.Position) *PositionResponse {
	if model == nil {
		return nil
	}

	return &PositionResponse{
		ID:         model.ID,
		Title:      model.Title,
		BaseSalary: model.BaseSalary,
	}
}

func NewPositionResponses(models []model.Position) []PositionResponse {
	result := make([]PositionResponse, 0, len(models))

	for i := range models {
		result = append(result, *NewPositionResponse(&models[i]))
	}
	return result
}
