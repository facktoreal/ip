package repositories

import "context"

// HealthRepository ...
type HealthRepository interface {
	Check(ctx context.Context) error
}
