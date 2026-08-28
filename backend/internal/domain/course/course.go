package course

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("course not found")

type Course struct {
	ID          int64
	Name        string
	Slug        string
	Description string
	Duration    string
	Level       string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}