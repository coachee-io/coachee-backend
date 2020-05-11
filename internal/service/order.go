package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"time"
)

// CreateOrder creates a new order
func (s *Service) CreateOrder(ctx context.Context, p *coachee.CreateOrderPayload) (res *coachee.CreateOrderResult, err error) {
	l := s.logger.With().Str("service", "CreateOrder").Logger()

	customerID, err := auth.GetCustomerID(p.Token)
	if err != nil {
		l.Error().Err(err).Msg("failed get id from jwt")
		return nil, err
	}

	customer, err := s.customerRepository.GetByID(repository.DefaultNoTransaction, customerID)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve customer")
		return nil, err
	}

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.CoachID)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve coach")
		return nil, err
	}

	program, err := coach.GetProgram(p.ProgramID)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve program")
		return nil, err
	}

	order := &model.Order{
		CoachID:    coach.ID,
		CustomerID: customer.ID,
		ProgramID:  program.ID,
		Amount:     program.TotalPrice,
		TaxPercent: 2000,
		Status:     model.OrderStatusCreated,
		IntroCall:  time.Unix(p.IntroCall, 0),
	}

	tx := s.orderRepository.Begin()
	if err := s.orderRepository.Create(tx, order); err != nil {
		l.Error().Err(err).Msg("failed store order")
		_ = tx.Rollback()
		return nil, err
	}

	secret, err := s.stripe.CreatePaymentIntent(order, customer, coach.StripeID)
	if err != nil {
		l.Error().Err(err).Msg("failed to create payment intent")
		_ = tx.Rollback()
		return nil, coachee.MakeTransient(err)
	}

	if err := s.orderRepository.Update(tx, order); err != nil {
		l.Error().Err(err).Msg("failed update order")
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		l.Error().Err(err).Msg("failed to commit db changes")
		return nil, coachee.MakeTransient(err)
	}

	return &coachee.CreateOrderResult{
		ClientSecret:  secret,
		PublishingKey: s.publishableKey,
	}, nil

}
