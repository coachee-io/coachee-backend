package service

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"

	"github.com/rs/zerolog"
)

// coachee service example implementation.
// The example methods log the requests and return zero values.
type Service struct {
	logger *zerolog.Logger

	coachRepository  repository.Coach
	clientRepository repository.Customer

	stripe Stripe
}

// Stripe is the interface for the stripe client
type Stripe interface {
	CreateCustomer(customer *model.Customer) error
}

// NewCoachee returns the coachee service implementation.
func NewCoachee(ctx context.Context, coach repository.Coach, client repository.Customer, stripe Stripe) *Service {
	log := zerolog.Ctx(ctx).With().Str("component", "service").Logger()
	return &Service{
		logger:           &log,
		coachRepository:  coach,
		clientRepository: client,
		stripe:           stripe,
	}
}
