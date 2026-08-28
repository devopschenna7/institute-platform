package course

import (
	"context"
	"fmt"
	"strings"

	domaincourse "institute-platform/backend/internal/domain/course"
)

type Service struct {
	repository domaincourse.Repository
}

func NewService(repository domaincourse.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetCourses(ctx context.Context) ([]domaincourse.Course, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) GetCourseBySlug(
	ctx context.Context,
	slug string,
) (*domaincourse.Course, error) {

	slug = strings.TrimSpace(slug)

	if slug == "" {
		return nil, fmt.Errorf("course slug is required")
	}

	return s.repository.GetBySlug(ctx, slug)
}