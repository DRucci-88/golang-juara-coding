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

func (dto *PositionResponse) FromModel(model *model.Position) {
	dto.ID = model.ID
	dto.Title = model.Title
	dto.BaseSalary = model.BaseSalary
}
