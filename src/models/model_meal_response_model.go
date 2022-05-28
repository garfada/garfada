package models

type MealResponseModel struct {

	Id int32 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Dishes []string `json:"dishes,omitempty"`
}
