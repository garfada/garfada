package models

import (
	"time"
)

type ServeResponseModel struct {

	Id int32 `json:"id,omitempty"`

	Meal int32 `json:"meal,omitempty"`

	Employee int32 `json:"employee,omitempty"`

	Date time.Time `json:"date,omitempty"`
}
