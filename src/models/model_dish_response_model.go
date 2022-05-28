package models

type DishResponseModel struct {

	Id int32 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`
}
