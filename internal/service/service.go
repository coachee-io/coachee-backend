package service

import (
	"coachee-backend/internal/repository"
	"context"

	"github.com/rs/zerolog"
)

// coachee service example implementation.
// The example methods log the requests and return zero values.
type Service struct {
	logger *zerolog.Logger

	coachRepository repository.Coach
}

// NewCoachee returns the coachee service implementation.
func NewCoachee(ctx context.Context, logger *zerolog.Logger, coach repository.Coach) *Service {
	return &Service{
		logger:          logger,
		coachRepository: coach,
	}
}
