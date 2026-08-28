package student

import "time"

type Student struct {
	ID          int64
	StudentCode string
	Name        string
	Email       string
	Phone       string
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
