package student

import "context"

type Repository interface {
	Create(ctx context.Context, student *Student) error
	GetByID(ctx context.Context, id int64) (*Student, error)
	GetByEmail(ctx context.Context, email string) (*Student, error)
	GetByStudentCode(
		ctx context.Context,
		studentCode string,
	) (*Student, error)
	GetAll(ctx context.Context) ([]Student, error)
	Update(ctx context.Context, student *Student) error
	GetDashboard(
		ctx context.Context,
		studentID int64,
	) (*Dashboard, error)
}
