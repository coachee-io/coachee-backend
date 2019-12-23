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

// CreateClient creates a new client and returns a jwt
func (s *Service) CreateClient(ctx context.Context, p *coachee.CreateClientPayload) (res *coachee.CreateClientResult, err error) {
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

	client := &model.Client{
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

	err = s.clientRepository.Create(repository.DefaultNoTransaction, client)
	if err != nil {
		l.Error().Err(err).Msg("failed to persist client")
		return nil, err
	}

	expiry := time.Now().Add(30 * time.Minute)
	token, err := auth.CreateUserToken(client.ID, expiry)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		return nil, coachee.MakeInternal(err)
	}

	return &coachee.CreateClientResult{
		Token: token,
		User: &coachee.BaseClient{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Expiry:    expiry.Unix(),
		},
	}, nil
}

// ClientLogin authenticates a client and returns a jwt
func (s *Service) ClientLogin(ctx context.Context, p *coachee.ClientLoginPayload) (*coachee.ClientLoginResult, error) {
	l := s.logger.With().Str("service", "ClientLogin").Logger()

	client, err := s.clientRepository.GetByEmail(repository.DefaultNoTransaction, p.Email)
	if err != nil {
		l.Debug().Err(err).Msg("failed to retrieve client")
		return nil, err
	}

	err = auth.VerifyPassword(client.Password, p.Password)
	if err != nil {
		l.Debug().Err(err).Msg("failed to authenticate")
		return nil, coachee.MakeValidation(errors.New("wrong password"))
	}

	expiry := time.Now().Add(30 * time.Minute)
	token, err := auth.CreateUserToken(client.ID, expiry)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		return nil, coachee.MakeInternal(err)
	}

	return &coachee.ClientLoginResult{
		Token: token,
		User: &coachee.BaseClient{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Expiry:    expiry.Unix(),
		},
	}, nil
}
