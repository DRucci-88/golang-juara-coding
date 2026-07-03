package dto

import (
	"time"
	"ujian2_rematch/model"
)

type AttendanceResponse struct {
	ID uint `json:"id"`

	Date     time.Time              `json:"date"`
	CheckIn  *time.Time             `json:"check_in,omitempty"`
	CheckOut *time.Time             `json:"check_out,omitempty"`
	Status   model.AttendanceStatus `json:"status"`

	EmployeeID uint              `json:"employee_id"`
	Employee   *EmployeeResponse `json:"employee,omitempty"`
}

func NewAttendanceResponse(model *model.Attendance) *AttendanceResponse {
	if model == nil {
		return nil
	}

	var checkIn *time.Time
	if model.CheckIn.Valid {
		checkIn = &model.CheckIn.Time
	}

	var checkOut *time.Time
	if model.CheckOut.Valid {
		checkOut = &model.CheckOut.Time
	}

	return &AttendanceResponse{
		ID: model.ID,

		Date:     model.Date,
		CheckIn:  checkIn,
		CheckOut: checkOut,
		Status:   model.Status,

		EmployeeID: model.EmployeeID,
		Employee:   NewEmployeeResponse(model.Employee),
	}
}

func NewAttendanceResponses(models []model.Attendance) []AttendanceResponse {
	result := make([]AttendanceResponse, 0, len(models))

	for i := range models {
		result = append(result, *NewAttendanceResponse(&models[i]))
	}

	return result
}
