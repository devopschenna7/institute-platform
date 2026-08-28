package auth

import "context"

type Repository interface {
	CreateAccount(
		ctx context.Context,
		account *Account,
	) error

	GetAccountByEmail(
		ctx context.Context,
		email string,
	) (*Account, error)

	GetAccountByStudentID(
		ctx context.Context,
		studentID int64,
	) (*Account, error)

	CreateSession(
		ctx context.Context,
		studentID int64,
	) (string, error)

	GetSessionStudentID(
		ctx context.Context,
		sessionID string,
	) (int64, error)

	DeleteSession(
		ctx context.Context,
		sessionID string,
	) error
}
