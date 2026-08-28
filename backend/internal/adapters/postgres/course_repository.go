package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domaincourse "institute-platform/backend/internal/domain/course"
)

type CourseRepository struct {
	db *pgxpool.Pool
}

func NewCourseRepository(db *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) GetAll(
	ctx context.Context,
) ([]domaincourse.Course, error) {

	const query = `
		SELECT
			id,
			name,
			slug,
			description,
			duration,
			level,
			status,
			created_at,
			updated_at
		FROM courses
		WHERE status = 'active'
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query courses: %w", err)
	}
	defer rows.Close()

	courses := make([]domaincourse.Course, 0)

	for rows.Next() {
		var c domaincourse.Course

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.Description,
			&c.Duration,
			&c.Level,
			&c.Status,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan course: %w", err)
		}

		courses = append(courses, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate courses: %w", err)
	}

	return courses, nil
}

func (r *CourseRepository) GetBySlug(
	ctx context.Context,
	slug string,
) (*domaincourse.Course, error) {

	const query = `
		SELECT
			id,
			name,
			slug,
			description,
			duration,
			level,
			status,
			created_at,
			updated_at
		FROM courses
		WHERE slug = $1
		  AND status = 'active'
	`

	var c domaincourse.Course

	err := r.db.QueryRow(ctx, query, slug).Scan(
		&c.ID,
		&c.Name,
		&c.Slug,
		&c.Description,
		&c.Duration,
		&c.Level,
		&c.Status,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domaincourse.ErrNotFound
		}

		return nil, fmt.Errorf("get course by slug: %w", err)
	}

	return &c, nil
}