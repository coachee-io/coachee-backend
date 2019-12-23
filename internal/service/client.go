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
func (s *Service) CreateClient(ctx context.Context, p *coachee.CreateClientPayload) (string, error) {
	l := s.logger.With().Str("service", "CreateClient").Logger()

	if len(p.Password) < 8 {
		msg := "password needs to be longer than 8 characters"
		l.Debug().Msg(msg)
		return "", coachee.MakeValidation(errors.New(msg))
	}

	hashedPass, err := auth.Hash(p.Password)
	if err != nil {
		l.Info().Err(err).Msg("failed to hash password")
		return "", coachee.MakeValidation(err)
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
		return "", coachee.MakeValidation(err)
	}

	err = s.clientRepository.Create(repository.DefaultNoTransaction, client)
	if err != nil {
		l.Error().Err(err).Msg("failed to persist client")
		return "", err
	}

	token, err := auth.CreateUserToken(client.ID)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		return "", coachee.MakeInternal(err)
	}

	return token, nil
}

// ClientLogin authenticates a client and returns a jwt
func (s *Service) ClientLogin(ctx context.Context, p *coachee.ClientLoginPayload) (string, error) {
	l := s.logger.With().Str("service", "ClientLogin").Logger()

	client, err := s.clientRepository.GetByEmail(repository.DefaultNoTransaction, p.Email)
	if err != nil {
		l.Debug().Err(err).Msg("failed to retrieve client")
		return "", err
	}

	err = auth.VerifyPassword(client.Password, p.Password)
	if err != nil {
		l.Debug().Err(err).Msg("failed to authenticate")
		return "", coachee.MakeValidation(err)
	}

	token, err := auth.CreateUserToken(client.ID)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		return "", coachee.MakeInternal(err)
	}

	return token, nil
}
