package mock

import (
	"context"

	"github.com/facktoreal/ip/lib/repositories"
)

// NewHealthRepository ...
func NewHealthRepository() repositories.HealthRepository {
	return &healthRepository{}
}

type healthRepository struct {
}

// Check ...
func (r *healthRepository) Check(ctx context.Context) error {
	return nil
}
