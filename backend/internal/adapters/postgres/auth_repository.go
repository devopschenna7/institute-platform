package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domainauth "institute-platform/backend/internal/domain/auth"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(
	db *pgxpool.Pool,
) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateAccount(
	ctx context.Context,
	account *domainauth.Account,
) error {
	query := `
		INSERT INTO student_accounts (
			student_id,
			email,
			password_hash
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			created_at,
			updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		account.StudentID,
		account.Email,
		account.PasswordHash,
	).Scan(
		&account.ID,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("create account: %w", err)
	}

	return nil
}

func (r *AuthRepository) GetAccountByEmail(
	ctx context.Context,
	email string,
) (*domainauth.Account, error) {
	query := `
		SELECT
			id,
			student_id,
			email,
			password_hash,
			created_at,
			updated_at
		FROM student_accounts
		WHERE email = $1
	`

	account := &domainauth.Account{}

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&account.ID,
		&account.StudentID,
		&account.Email,
		&account.PasswordHash,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get account by email: %w", err)
	}

	return account, nil
}

func (r *AuthRepository) GetAccountByStudentID(
	ctx context.Context,
	studentID int64,
) (*domainauth.Account, error) {
	query := `
		SELECT
			id,
			student_id,
			email,
			password_hash,
			created_at,
			updated_at
		FROM student_accounts
		WHERE student_id = $1
	`

	account := &domainauth.Account{}

	err := r.db.QueryRow(
		ctx,
		query,
		studentID,
	).Scan(
		&account.ID,
		&account.StudentID,
		&account.Email,
		&account.PasswordHash,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf(
			"get account by student id: %w",
			err,
		)
	}

	return account, nil
}

func (r *AuthRepository) CreateSession(
	ctx context.Context,
	studentID int64,
) (string, error) {
	sessionID := uuid.New().String()

	query := `
		INSERT INTO sessions (
			id,
			student_id,
			expires_at
		)
		VALUES ($1, $2, $3)
	`

	expiresAt := time.Now().Add(24 * time.Hour)

	_, err := r.db.Exec(
		ctx,
		query,
		sessionID,
		studentID,
		expiresAt,
	)

	if err != nil {
		return "", fmt.Errorf(
			"create session: %w",
			err,
		)
	}

	return sessionID, nil
}

func (r *AuthRepository) GetSessionStudentID(
	ctx context.Context,
	sessionID string,
) (int64, error) {
	query := `
		SELECT student_id
		FROM sessions
		WHERE id = $1
		  AND expires_at > CURRENT_TIMESTAMP
	`

	var studentID int64

	err := r.db.QueryRow(
		ctx,
		query,
		sessionID,
	).Scan(&studentID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("session not found")
		}

		return 0, fmt.Errorf(
			"get session: %w",
			err,
		)
	}

	return studentID, nil
}

func (r *AuthRepository) DeleteSession(
	ctx context.Context,
	sessionID string,
) error {
	query := `
		DELETE FROM sessions
		WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		sessionID,
	)

	if err != nil {
		return fmt.Errorf(
			"delete session: %w",
			err,
		)
	}

	return nil
}
