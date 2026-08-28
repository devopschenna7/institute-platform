package course

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Course, error)
	GetBySlug(ctx context.Context, slug string) (*Course, error)
}