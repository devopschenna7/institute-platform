package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domainregistration "institute-platform/backend/internal/domain/registration"
)

type RegistrationRepository struct {
	db *pgxpool.Pool
}

func NewRegistrationRepository(
	db *pgxpool.Pool,
) *RegistrationRepository {
	return &RegistrationRepository{
		db: db,
	}
}

func (r *RegistrationRepository) Create(
	ctx context.Context,
	registration *domainregistration.Registration,
) error {
	query := `
		INSERT INTO registrations (
			student_id,
			course_id,
			status
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
		registration.StudentID,
		registration.CourseID,
		registration.Status,
	).Scan(
		&registration.ID,
		&registration.CreatedAt,
		&registration.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("create registration: %w", err)
	}

	return nil
}

func (r *RegistrationRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domainregistration.Registration, error) {
	query := `
		SELECT
			id,
			student_id,
			course_id,
			status,
			created_at,
			updated_at
		FROM registrations
		WHERE id = $1
	`

	registration := &domainregistration.Registration{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&registration.ID,
		&registration.StudentID,
		&registration.CourseID,
		&registration.Status,
		&registration.CreatedAt,
		&registration.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("registration not found: %w", err)
		}

		return nil, fmt.Errorf("get registration: %w", err)
	}

	return registration, nil
}

func (r *RegistrationRepository) GetByStudentID(
	ctx context.Context,
	studentID int64,
) ([]domainregistration.Registration, error) {
	query := `
		SELECT
			id,
			student_id,
			course_id,
			status,
			created_at,
			updated_at
		FROM registrations
		WHERE student_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, fmt.Errorf(
			"get registrations by student: %w",
			err,
		)
	}
	defer rows.Close()

	registrations := make(
		[]domainregistration.Registration,
		0,
	)

	for rows.Next() {
		var registration domainregistration.Registration

		if err := rows.Scan(
			&registration.ID,
			&registration.StudentID,
			&registration.CourseID,
			&registration.Status,
			&registration.CreatedAt,
			&registration.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf(
				"scan registration: %w",
				err,
			)
		}

		registrations = append(
			registrations,
			registration,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate registrations: %w",
			err,
		)
	}

	return registrations, nil
}

func (r *RegistrationRepository) GetByCourseID(
	ctx context.Context,
	courseID int64,
) ([]domainregistration.Registration, error) {
	query := `
		SELECT
			id,
			student_id,
			course_id,
			status,
			created_at,
			updated_at
		FROM registrations
		WHERE course_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query, courseID)
	if err != nil {
		return nil, fmt.Errorf(
			"get registrations by course: %w",
			err,
		)
	}
	defer rows.Close()

	registrations := make(
		[]domainregistration.Registration,
		0,
	)

	for rows.Next() {
		var registration domainregistration.Registration

		if err := rows.Scan(
			&registration.ID,
			&registration.StudentID,
			&registration.CourseID,
			&registration.Status,
			&registration.CreatedAt,
			&registration.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf(
				"scan registration: %w",
				err,
			)
		}

		registrations = append(
			registrations,
			registration,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate registrations: %w",
			err,
		)
	}

	return registrations, nil
}

func (r *RegistrationRepository) GetByStudentAndCourse(
	ctx context.Context,
	studentID int64,
	courseID int64,
) (*domainregistration.Registration, error) {
	query := `
		SELECT
			id,
			student_id,
			course_id,
			status,
			created_at,
			updated_at
		FROM registrations
		WHERE student_id = $1
		  AND course_id = $2
	`

	registration := &domainregistration.Registration{}

	err := r.db.QueryRow(
		ctx,
		query,
		studentID,
		courseID,
	).Scan(
		&registration.ID,
		&registration.StudentID,
		&registration.CourseID,
		&registration.Status,
		&registration.CreatedAt,
		&registration.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf(
			"check existing registration: %w",
			err,
		)
	}

	return registration, nil
}

func (r *RegistrationRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status domainregistration.Status,
) error {
	query := `
		UPDATE registrations
		SET
			status = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	result, err := r.db.Exec(
		ctx,
		query,
		status,
		id,
	)
	if err != nil {
		return fmt.Errorf(
			"update registration status: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("registration not found")
	}

	return nil
}
