package dto

import (
	"time"
	"ujian2_rematch/model"

	"github.com/go-playground/validator/v10"
)

type LeaveCreateRequest struct {
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Reason    string    `json:"reason" binding:"required,max=255"`
}

type LeaveApprovalRequest struct {
	Status model.LeaveStatus `json:"status"`
}

func LeaveRequestValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(LeaveCreateRequest)

	if req.EndDate.Before(req.StartDate) {
		sl.ReportError(
			req.EndDate,
			"EndDate",
			"end_date",
			"afterstart",
			"",
		)
	}
}

type LeaveResponse struct {
	ID         uint              `json:"id"`
	EmployeeID uint              `json:"employee_id"`
	Employee   *EmployeeResponse `json:"employee,omitzero"`
	StartDate  time.Time         `json:"start_date"`
	EndDate    time.Time         `json:"end_date"`
	Reason     string            `json:"reason"`
	Status     model.LeaveStatus `json:"status"`
}

func NewLeaveResponse(model *model.Leave) *LeaveResponse {
	if model == nil {
		return nil
	}

	return &LeaveResponse{
		ID:         model.ID,
		EmployeeID: model.EmployeeID,
		Employee:   NewEmployeeResponse(model.Employee),
		StartDate:  model.StartDate,
		EndDate:    model.EndDate,
		Reason:     model.Reason,
		Status:     model.Status,
	}
}

func NewLeaveResponses(models []model.Leave) []LeaveResponse {
	responses := make([]LeaveResponse, 0, len(models))

	for i := range models {
		responses = append(
			responses,
			*NewLeaveResponse(&models[i]),
		)
	}

	return responses
}
