package models

type EmployeeResponseModel struct {

	Id int32 `json:"id,omitempty"`

	Code string `json:"code,omitempty"`

	Name string `json:"name,omitempty"`

	Business int32 `json:"business,omitempty"`
}
