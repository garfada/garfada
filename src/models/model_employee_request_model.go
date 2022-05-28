package models

type EmployeeRequestModel struct {

	Code string `json:"code"`

	Name string `json:"name"`

	Business int32 `json:"business"`
}
