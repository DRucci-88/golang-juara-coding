package dto

import (
	"ujian2_rematch/model"
)

type DepartmentCreateRequest struct {
	Name string `json:"name" binding:"required,max=50"`
	Code string `json:"code" binding:"required,max=50"`
}

type DepartmentUpdateRequest struct {
	Name string `json:"name" binding:"required,max=50"`
	Code string `json:"code" binding:"max=50"`
}

type DepartmentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewDepartmentResponse(model *model.Department) *DepartmentResponse {
	if model == nil {
		return nil
	}

	return &DepartmentResponse{
		ID:   model.ID,
		Name: model.Name,
		Code: model.Code,
	}
}

func NewDepartmentResponses(models []model.Department) []DepartmentResponse {
	result := make([]DepartmentResponse, 0, len(models))

	for i := range models {
		result = append(result, *NewDepartmentResponse(&models[i]))
	}
	return result
}
