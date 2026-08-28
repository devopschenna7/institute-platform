package registration

import (
	"context"
	"fmt"

	domainregistration "institute-platform/backend/internal/domain/registration"
)

type Service struct {
	repository domainregistration.Repository
}

func NewService(
	repository domainregistration.Repository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateRegistration(
	ctx context.Context,
	studentID int64,
	courseID int64,
) (*domainregistration.Registration, error) {

	existing, err := s.repository.GetByStudentAndCourse(
		ctx,
		studentID,
		courseID,
	)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, fmt.Errorf(
			"student is already registered for this course",
		)
	}

	registration := &domainregistration.Registration{
		StudentID: studentID,
		CourseID:  courseID,
		Status:    domainregistration.StatusPending,
	}

	if err := s.repository.Create(
		ctx,
		registration,
	); err != nil {
		return nil, err
	}

	return registration, nil
}

func (s *Service) GetRegistration(
	ctx context.Context,
	id int64,
) (*domainregistration.Registration, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetStudentRegistrations(
	ctx context.Context,
	studentID int64,
) ([]domainregistration.Registration, error) {
	return s.repository.GetByStudentID(
		ctx,
		studentID,
	)
}

func (s *Service) GetCourseRegistrations(
	ctx context.Context,
	courseID int64,
) ([]domainregistration.Registration, error) {
	return s.repository.GetByCourseID(
		ctx,
		courseID,
	)
}

func (s *Service) UpdateStatus(
	ctx context.Context,
	id int64,
	status domainregistration.Status,
) error {
	return s.repository.UpdateStatus(
		ctx,
		id,
		status,
	)
}
