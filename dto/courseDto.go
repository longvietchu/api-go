package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateCourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	Level       string `json:"level"`
	URL         string `json:"url"`
	Language    string `json:"language"`
	Commitment  string `json:"commitment"`
	Rating      string `json:"rating"`
}

func (c CreateCourseRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required.Error("Thông tin này là bắt buộc!"), validation.Length(5, 100)),
	)
}
