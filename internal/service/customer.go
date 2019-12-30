package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"errors"
	"time"
)

// CreateCustomer creates a new customer and returns a jwt
func (s *Service) CreateCustomer(ctx context.Context, p *coachee.CreateCustomerPayload) (res *coachee.CreateCustomerResult, err error) {
	l := s.logger.With().Str("service", "CreateClient").Logger()

	if len(p.Password) < 8 {
		msg := "password needs to be longer than 8 characters"
		l.Debug().Msg(msg)
		return nil, coachee.MakeValidation(errors.New(msg))
	}

	hashedPass, err := auth.Hash(p.Password)
	if err != nil {
		l.Info().Err(err).Msg("failed to hash password")
		return nil, coachee.MakeValidation(err)
	}

	client := &model.Customer{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		BirthDate: time.Unix(p.BirthDate, 0),
		Password:  string(hashedPass),
	}

	err = client.Validate()
	if err != nil {
		l.Debug().Err(err).Msg("client validation failed")
		return nil, coachee.MakeValidation(err)
	}

	tx := s.customerRepository.Begin()
	err = s.customerRepository.Create(tx, client)
	if err != nil {
		l.Error().Err(err).Msg("failed to persist client")
		_ = tx.Rollback()
		return nil, err
	}

	err = s.stripe.CreateCustomer(client)
	if err != nil {
		l.Error().Err(err).Msg("failed to store client in stripe")
		_ = tx.Rollback()
		return nil, coachee.MakeTransient(err)
	}

	err = s.customerRepository.Update(tx, client)
	if err != nil {
		l.Error().Err(err).Msg("failed to update client")
		_ = tx.Rollback()
		return nil, coachee.MakeTransient(err)
	}

	createdAt := time.Now()
	token, err := auth.CreateUserToken(client.ID, createdAt)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		_ = tx.Rollback()
		return nil, coachee.MakeInternal(err)
	}

	if err := tx.Commit(); err != nil {
		l.Error().Err(err).Msg("failed to commit client changes")
		return nil, coachee.MakeTransient(err)
	}

	return &coachee.CreateCustomerResult{
		Token:  token,
		Expiry: createdAt.Add(30 * time.Minute).Unix(),
		User: &coachee.BaseClient{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
		},
	}, nil
}

// ClientLogin authenticates a customer and returns a jwt
func (s *Service) CustomerLogin(ctx context.Context, p *coachee.CustomerLoginPayload) (res *coachee.CustomerLoginResult, err error) {
	l := s.logger.With().Str("service", "ClientLogin").Logger()

	client, err := s.customerRepository.GetByEmail(repository.DefaultNoTransaction, p.Email)
	if err != nil {
		l.Debug().Err(err).Msg("failed to retrieve client")
		return nil, err
	}

	err = auth.VerifyPassword(client.Password, p.Password)
	if err != nil {
		l.Debug().Err(err).Msg("failed to authenticate")
		return nil, coachee.MakeValidation(errors.New("wrong password"))
	}

	createdAt := time.Now()
	token, err := auth.CreateUserToken(client.ID, createdAt)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		return nil, coachee.MakeInternal(err)
	}

	return &coachee.CustomerLoginResult{
		Token:  token,
		Expiry: createdAt.Add(30 * time.Minute).Unix(),
		User: &coachee.BaseClient{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
		},
	}, nil
}
