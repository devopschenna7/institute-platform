package registration

import "context"

type Repository interface {
	Create(ctx context.Context, registration *Registration) error

	GetByID(
		ctx context.Context,
		id int64,
	) (*Registration, error)

	GetByStudentID(
		ctx context.Context,
		studentID int64,
	) ([]Registration, error)

	GetByCourseID(
		ctx context.Context,
		courseID int64,
	) ([]Registration, error)

	GetByStudentAndCourse(
		ctx context.Context,
		studentID int64,
		courseID int64,
	) (*Registration, error)

	UpdateStatus(
		ctx context.Context,
		id int64,
		status Status,
	) error
}
