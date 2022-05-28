package models

type MealRequestModel struct {

	Name string `json:"name"`

	Description string `json:"description"`

	Dishes []string `json:"dishes"`
}
