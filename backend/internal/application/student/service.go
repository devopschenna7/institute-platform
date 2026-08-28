package student

import (
	"context"

	domainstudent "institute-platform/backend/internal/domain/student"
)

type Service struct {
	repository domainstudent.Repository
}

func NewService(
	repository domainstudent.Repository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateStudent(
	ctx context.Context,
	student *domainstudent.Student,
) error {
	return s.repository.Create(ctx, student)
}

func (s *Service) GetStudent(
	ctx context.Context,
	id int64,
) (*domainstudent.Student, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetStudents(
	ctx context.Context,
) ([]domainstudent.Student, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) UpdateStudent(
	ctx context.Context,
	student *domainstudent.Student,
) error {
	return s.repository.Update(ctx, student)
}

func (s *Service) GetDashboard(
	ctx context.Context,
	studentID int64,
) (*domainstudent.Dashboard, error) {
	return s.repository.GetDashboard(
		ctx,
		studentID,
	)
}
