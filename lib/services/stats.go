package services

import (
	"context"
	"time"

	"github.com/facktoreal/ip/lib/models"
)

type statsService struct {
	data models.Stats
}

// StatsService ...
type StatsService interface {
	Get(ctx context.Context) models.Stats
}

// NewStatsService ...
func NewStatsService() StatsService {
	return &statsService{
		models.Stats{Uptime: time.Now()},
	}
}

// Get ...
func (s *statsService) Get(ctx context.Context) models.Stats {
	return s.data
}
