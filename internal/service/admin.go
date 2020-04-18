package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"context"
	"errors"
	"time"
)

// AdminLogin validates an admin login and returns the respective token
func (s *Service) AdminLogin(ctx context.Context, p *coachee.AdminLoginPayload) (*coachee.AdminLoginResult, error) {
	l := s.logger.With().Str("service", "AdminLogin").Logger()

	if p.Email != s.adminEmail {
		l.Debug().Msg("wrong admin email")
		return nil, coachee.MakeValidation(errors.New("wrong login details"))
	}

	err := auth.VerifyPassword(s.password, p.Password)
	if err != nil {
		l.Debug().Err(err).Msg("wrong password")
		return nil, coachee.MakeValidation(errors.New("wrong login details"))
	}

	createdAt := time.Now()
	token, err := auth.CreateAdminToken(createdAt)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate jwt")
		return nil, coachee.MakeInternal(err)
	}

	return &coachee.AdminLoginResult{
		Token:  token,
		Expiry: createdAt.Add(60 * time.Minute).Unix(),
	}, nil
}
