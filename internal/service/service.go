package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
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
	newsletterRepository    repository.Newsletter

	stripe Stripe
	email  Email
	slack  Slack

	publishableKey string
	adminEmail     string
	password       string
}

// Stripe is the interface for the stripe client
type Stripe interface {
	CreateCustomer(customer *model.Customer) error
	CreatePaymentIntent(order *model.Order, customer *model.Customer, coachID string) (string, error)
	RegisterStripeExpress(authCode string) (string, error)
	LoginStripeExpress(stripeID string) (string, error)
}

// Email is the email client to send emails
type Email interface {
	SendBookingEmail(to, programme, coachName string) error
	SendWelcomeEmail(to, token string) error
	SendClientPasswordRecoveryEmail(to, token string) error
	SendCoachPasswordRecoveryEmail(to, token string) error
}

type Slack interface {
	Post(message []byte) error
}

type Config struct {
	Coach         repository.Coach
	Customer      repository.Customer
	Order         repository.Order
	Recovery      repository.Recovery
	CoachRecovery repository.CoachRecovery
	Newsletter    repository.Newsletter
	Stripe        Stripe
	Email         Email
	Slack         Slack
	PubKey        string
	AdminEmail    string
	AdminPassword string
}

// NewCoachee returns the coachee service implementation.
func NewCoachee(config Config) (*Service, error) {

	log := zerolog.New(os.Stderr).With().Timestamp().Str("component", "service").Logger()

	hashedPass, err := auth.Hash(config.AdminPassword)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash admin password")
		return nil, coachee.MakeValidation(err)
	}

	return &Service{
		logger:                  &log,
		coachRepository:         config.Coach,
		customerRepository:      config.Customer,
		orderRepository:         config.Order,
		recoveryRepository:      config.Recovery,
		coachRecoveryRepository: config.CoachRecovery,
		newsletterRepository:    config.Newsletter,
		stripe:                  config.Stripe,
		email:                   config.Email,
		slack:                   config.Slack,
		publishableKey:          config.PubKey,
		adminEmail:              config.AdminEmail,
		password:                string(hashedPass),
	}, nil
}
