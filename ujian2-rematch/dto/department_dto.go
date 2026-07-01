package dto

import "ujian2_rematch/model"

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

func (dto *DepartmentResponse) FromModel(model *model.Department) {
	dto.ID = model.ID
	dto.Name = model.Name
	dto.Code = model.Code
}
