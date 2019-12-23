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

	coachRepository  repository.Coach
	clientRepository repository.Client
}

// NewCoachee returns the coachee service implementation.
func NewCoachee(ctx context.Context, coach repository.Coach, client repository.Client) *Service {
	log := zerolog.Ctx(ctx).With().Str("component", "service").Logger()
	return &Service{
		logger:           &log,
		coachRepository:  coach,
		clientRepository: client,
	}
}
