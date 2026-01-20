package services

import (
	"context"

	"github.com/facktoreal/ip/app/lib/repositories"
)

type healthService struct {
	repo repositories.HealthRepository
}

// HealthService ...
type HealthService interface {
	Check(ctx context.Context) error
}

// NewHealthService ...
func NewHealthService(repo repositories.HealthRepository) HealthService {
	return &healthService{
		repo: repo,
	}
}

// Check ...
func (s *healthService) Check(ctx context.Context) error {
	return s.repo.Check(ctx)
}
