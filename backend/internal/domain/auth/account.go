package auth

import "time"

type Account struct {
	ID           int64
	StudentID    int64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
