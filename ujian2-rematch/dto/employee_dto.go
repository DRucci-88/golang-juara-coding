package dto

import "ujian2_rematch/model"

type EmployeeCreateRequest struct {
	NIK      string `json:"nik" binding:"required,max=50"`
	FullName string `json:"full_name" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`

	DepartmentID uint `json:"department_id" binding:"required"`
	PositionID   uint `json:"position_id" binding:"required"`
}

type EmployeeResponse struct {
	ID       uint   `json:"id"`
	NIK      string `json:"nik"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`

	DepartmentID uint                `json:"department_id"`
	Department   *DepartmentResponse `json:"department,omitempty"`
	PositionID   uint                `json:"position_id"`
	Position     *PositionResponse   `json:"position,omitempty"`
}

func NewEmployeeResponse(model *model.Employee) *EmployeeResponse {
	if model == nil {
		return nil
	}

	return &EmployeeResponse{
		ID:       model.ID,
		NIK:      model.NIK,
		FullName: model.FullName,
		Email:    model.Email,

		DepartmentID: model.DepartmentID,
		Department:   NewDepartmentResponse(model.Department),
		PositionID:   model.PositionID,
		Position:     NewPositionResponse(model.Position),
	}
}

func NewEmployeeResponses(models []model.Employee) []EmployeeResponse {
	responses := make([]EmployeeResponse, 0, len(models))

	for i := range models {
		responses = append(
			responses,
			*NewEmployeeResponse(&models[i]),
		)
	}

	return responses
}
