package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domainstudent "institute-platform/backend/internal/domain/student"
)

type StudentRepository struct {
	db *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) Create(
	ctx context.Context,
	student *domainstudent.Student,
) error {
	query := `
		INSERT INTO students (
			student_code,
			name,
			email,
			phone,
			date_of_birth
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			created_at,
			updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		student.StudentCode,
		student.Name,
		student.Email,
		student.Phone,
		student.DateOfBirth,
	).Scan(
		&student.ID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("create student: %w", err)
	}

	return nil
}

func (r *StudentRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domainstudent.Student, error) {
	query := `
			SELECT
				id,
				name,
				email,
				phone,
				date_of_birth,
				created_at,
				updated_at
			FROM students
			WHERE id = $1
		`

	student := &domainstudent.Student{}

	err := r.db.QueryRow(ctx, query, id).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Phone,
		&student.DateOfBirth,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("student not found: %w", err)
		}

		return nil, fmt.Errorf("get student by id: %w", err)
	}

	return student, nil
}

func (r *StudentRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*domainstudent.Student, error) {
	query := `
			SELECT
				id,
				name,
				email,
				phone,
				date_of_birth,
				created_at,
				updated_at
			FROM students
			WHERE email = $1
		`

	student := &domainstudent.Student{}

	err := r.db.QueryRow(ctx, query, email).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Phone,
		&student.DateOfBirth,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("student not found: %w", err)
		}

		return nil, fmt.Errorf("get student by email: %w", err)
	}

	return student, nil
}

func (r *StudentRepository) GetAll(
	ctx context.Context,
) ([]domainstudent.Student, error) {
	query := `
			SELECT
				id,
				name,
				email,
				phone,
				date_of_birth,
				created_at,
				updated_at
			FROM students
			ORDER BY id
		`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get students: %w", err)
	}
	defer rows.Close()

	students := make([]domainstudent.Student, 0)

	for rows.Next() {
		var student domainstudent.Student

		if err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Email,
			&student.Phone,
			&student.DateOfBirth,
			&student.CreatedAt,
			&student.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan student: %w", err)
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate students: %w", err)
	}

	return students, nil
}

func (r *StudentRepository) Update(
	ctx context.Context,
	student *domainstudent.Student,
) error {
	query := `
			UPDATE students
			SET
				name = $1,
				email = $2,
				phone = $3,
				date_of_birth = $4,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $5
			RETURNING updated_at
		`

	err := r.db.QueryRow(
		ctx,
		query,
		student.Name,
		student.Email,
		student.Phone,
		student.DateOfBirth,
		student.ID,
	).Scan(&student.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("student not found: %w", err)
		}

		return fmt.Errorf("update student: %w", err)
	}

	return nil
}

func (r *StudentRepository) GetDashboard(
	ctx context.Context,
	studentID int64,
) (*domainstudent.Dashboard, error) {

	studentQuery := `
			SELECT
				id,
				name,
				email,
				phone,
				date_of_birth,
				created_at,
				updated_at
			FROM students
			WHERE id = $1
		`

	var dashboard domainstudent.Dashboard

	err := r.db.QueryRow(
		ctx,
		studentQuery,
		studentID,
	).Scan(
		&dashboard.Student.ID,
		&dashboard.Student.Name,
		&dashboard.Student.Email,
		&dashboard.Student.Phone,
		&dashboard.Student.DateOfBirth,
		&dashboard.Student.CreatedAt,
		&dashboard.Student.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("student not found")
		}

		return nil, fmt.Errorf(
			"get dashboard student: %w",
			err,
		)
	}

	registrationQuery := `
			SELECT
				r.id,
				r.course_id,
				c.name,
				c.slug,
				c.duration,
				c.level,
				r.status,
				r.created_at
			FROM registrations r
			INNER JOIN courses c
				ON c.id = r.course_id
			WHERE r.student_id = $1
			ORDER BY r.created_at DESC
		`

	rows, err := r.db.Query(
		ctx,
		registrationQuery,
		studentID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get dashboard registrations: %w",
			err,
		)
	}
	defer rows.Close()

	dashboard.Registrations = make(
		[]domainstudent.DashboardRegistration,
		0,
	)

	for rows.Next() {
		var registration domainstudent.DashboardRegistration

		if err := rows.Scan(
			&registration.RegistrationID,
			&registration.CourseID,
			&registration.CourseName,
			&registration.CourseSlug,
			&registration.CourseDuration,
			&registration.CourseLevel,
			&registration.Status,
			&registration.RegisteredAt,
		); err != nil {
			return nil, fmt.Errorf(
				"scan dashboard registration: %w",
				err,
			)
		}

		dashboard.Registrations = append(
			dashboard.Registrations,
			registration,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate dashboard registrations: %w",
			err,
		)
	}

	return &dashboard, nil
}

// 	func (r *StudentRepository) GetByStudentID(
//     ctx context.Context,
//     studentID string,
// ) (*domainstudent.Student, error) {

//     query := `
//         SELECT
//             id,
//             student_id,
//             name,
//             email,
//             phone,
//             date_of_birth,
//             created_at,
//             updated_at
//         FROM students
//         WHERE student_id = $1
//     `

//     student := &domainstudent.Student{}

//     err := r.db.QueryRow(
//         ctx,
//         query,
//         studentID,
//     ).Scan(
//         &student.ID,
//         &student.StudentID,
//         &student.Name,
//         &student.Email,
//         &student.Phone,
//         &student.DateOfBirth,
//         &student.CreatedAt,
//         &student.UpdatedAt,
//     )

//     if err != nil {
//         if errors.Is(err, pgx.ErrNoRows) {
//             return nil, nil
//         }

//         return nil, fmt.Errorf(
//             "get student by student id: %w",
//             err,
//         )
//     }

//     return student, nil
// }

func (r *StudentRepository) GetByStudentCode(
	ctx context.Context,
	studentCode string,
) (*domainstudent.Student, error) {

	query := `
		SELECT
			id,
			student_code,
			name,
			email,
			phone,
			date_of_birth,
			created_at,
			updated_at
		FROM students
		WHERE student_code = $1
	`

	student := &domainstudent.Student{}

	err := r.db.QueryRow(
		ctx,
		query,
		studentCode,
	).Scan(
		&student.ID,
		&student.StudentCode,
		&student.Name,
		&student.Email,
		&student.Phone,
		&student.DateOfBirth,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf(
			"get student by code: %w",
			err,
		)
	}

	return student, nil
}
