package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	domainauth "institute-platform/backend/internal/domain/auth"
	domainstudent "institute-platform/backend/internal/domain/student"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrAccountExists      = errors.New("account already exists")
	ErrStudentNotFound    = errors.New("student not found")
)

type Service struct {
	repository        domainauth.Repository
	studentRepository domainstudent.Repository
}

func NewService(
	repository domainauth.Repository,
	studentRepository domainstudent.Repository,
) *Service {
	return &Service{
		repository:        repository,
		studentRepository: studentRepository,
	}
}

func (s *Service) Register(
	ctx context.Context,
	studentCode string,
	email string,
	password string,
) error {
	studentCode = strings.TrimSpace(studentCode)
	email = strings.TrimSpace(strings.ToLower(email))

	if studentCode == "" ||
		email == "" ||
		password == "" {
		return errors.New(
			"student_code, email and password are required",
		)
	}

	// student, err := s.studentRepository.GetByStudentCode(
	// 	ctx,
	// 	studentCode,
	// )
	// if err != nil {
	// 	return err
	// }

	// if student == nil {
	// 	return ErrStudentNotFound
	// }

	student := &domainstudent.Student{
		StudentCode: studentCode,
		Name:        studentCode,
		Email:       email,
	}

	err := s.studentRepository.Create(ctx, student)
	if err != nil {
		return fmt.Errorf("create student: %w", err)
	}

	existing, err := s.repository.GetAccountByEmail(
		ctx,
		email,
	)
	if err != nil {
		return err
	}

	if existing != nil {
		return ErrAccountExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf(
			"hash password: %w",
			err,
		)
	}

	account := &domainauth.Account{
		StudentID:    student.ID,
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	return s.repository.CreateAccount(
		ctx,
		account,
	)
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	account, err := s.repository.GetAccountByEmail(
		ctx,
		email,
	)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if account == nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(account.PasswordHash),
		[]byte(password),
	); err != nil {
		return "", ErrInvalidCredentials
	}

	sessionID, err := s.repository.CreateSession(
		ctx,
		account.StudentID,
	)
	if err != nil {
		return "", fmt.Errorf("create session: %w", err)
	}

	return sessionID, nil
}

func (s *Service) Logout(
	ctx context.Context,
	sessionID string,
) error {
	return s.repository.DeleteSession(
		ctx,
		sessionID,
	)
}

func (s *Service) GetSessionStudentID(
	ctx context.Context,
	sessionID string,
) (int64, error) {
	return s.repository.GetSessionStudentID(
		ctx,
		sessionID,
	)
}
