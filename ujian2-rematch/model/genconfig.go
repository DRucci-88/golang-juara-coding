package model

import (
	"gorm.io/cli/gorm/field"
	"gorm.io/cli/gorm/genconfig"
)

var _ = genconfig.Config{
	FieldTypeMap: map[any]any{
		AttendanceStatus(""): field.String{},
		EmployeeStatus(""):   field.String{},
		LeaveStatus(""):      field.String{},
	},
	// FieldNameMap: map[string]any{
	// 	// "status": field.String{},
	// },
}
