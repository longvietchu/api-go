package entity

import (
	"errors"
	"time"
)

type Course struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Creator     string    `json:"creator,omitempty"`
	Level       string    `json:"level,omitempty"`
	URL         string    `json:"url,omitempty"`
	Language    string    `json:"language,omitempty"`
	Commitment  string    `json:"commitment,omitempty"`
	Rating      string    `json:"rating,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCourse(title, description, creator, level, url, language, commitment, rating string) (*Course, error) {
	if title == "" {
		return nil, errors.New("Title must not empty")
	}

	course := &Course{
		Title:       title,
		Description: description,
		Creator:     creator,
		Level:       level,
		URL:         url,
		Language:    language,
		Commitment:  commitment,
		Rating:      rating,
	}

	return course, nil
}

func (course *Course) SetCourse(title, description, creator, level, url, language, commitment, rating string) error {
	if title == "" {
		return errors.New("Title must not empty")
	}

	course.Title = title
	course.Description = description
	course.Creator = creator
	course.Level = level
	course.URL = url
	course.Language = language
	course.Commitment = commitment
	course.Rating = rating

	return nil
}
