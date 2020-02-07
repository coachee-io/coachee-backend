package service

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"os"

	"github.com/rs/zerolog"
)

// coachee service example implementation.
// The example methods log the requests and return zero values.
type Service struct {
	logger *zerolog.Logger

	coachRepository         repository.Coach
	customerRepository      repository.Customer
	orderRepository         repository.Order
	recoveryRepository      repository.Recovery
	coachRecoveryRepository repository.CoachRecovery

	stripe         Stripe
	publishableKey string
}

// Stripe is the interface for the stripe client
type Stripe interface {
	CreateCustomer(customer *model.Customer) error
	CreatePaymentIntent(order *model.Order, customer *model.Customer) (string, error)
	RegisterStripeExpress(authCode string) (string, error)
	LoginStripeExpress(stripeID string) (string, error)
}

// NewCoachee returns the coachee service implementation.
func NewCoachee(ctx context.Context,
	coach repository.Coach,
	client repository.Customer,
	order repository.Order,
	recovery repository.Recovery,
	coachRecovery repository.CoachRecovery,
	stripe Stripe,
	pubKey string) *Service {

	log := zerolog.New(os.Stderr).With().Timestamp().Str("component", "service").Logger()
	return &Service{
		logger:                  &log,
		coachRepository:         coach,
		customerRepository:      client,
		orderRepository:         order,
		recoveryRepository:      recovery,
		coachRecoveryRepository: coachRecovery,
		stripe:                  stripe,
		publishableKey:          pubKey,
	}
}
