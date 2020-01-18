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

// StartPasswordRecoveryFlow start the password recovery process
func (s *Service) StartPasswordRecoveryFlow(ctx context.Context, p *coachee.StartPasswordRecoveryFlowPayload) error {
	l := s.logger.With().Str("service", "StartPasswordRecoveryFlow").Str("email", p.Email).Logger()

	customer, err := s.customerRepository.GetByEmail(repository.DefaultNoTransaction, p.Email)
	if err != nil {
		l.Info().Err(err).Msg("failed to retrieve customer")
		return err
	}

	recovery := &model.Recovery{
		CustomerID: customer.ID,
	}

	if err := s.recoveryRepository.Create(repository.DefaultNoTransaction, recovery); err != nil {
		l.Error().Err(err).Msg("failed to persist recovery")
		return err
	}

	// TODO: send email, check db for now
	return nil
}

// CheckPasswordRecoveryToken checks if the recovery token is still valid
func (s *Service) CheckPasswordRecoveryToken(ctx context.Context, p *coachee.CheckPasswordRecoveryTokenPayload) error {
	l := s.logger.With().Str("service", "CheckPasswordRecoveryToken").Str("token", p.Token).Logger()

	recovery, err := s.recoveryRepository.GetByID(repository.DefaultNoTransaction, p.Token)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve recovery")
		return err
	}

	if !recoveryIsValid(recovery) {
		msg := "recovery token is expired"
		l.Info().Msg(msg)
		return coachee.MakeUnauthorized(errors.New(msg))
	}

	return nil
}

// FinalizePasswordRecoveryFlow checks if the token is valid and resets the password accordingly
func (s *Service) FinalizePasswordRecoveryFlow(ctx context.Context, p *coachee.FinalizePasswordRecoveryFlowPayload) error {
	l := s.logger.With().Str("service", "CheckPasswordRecoveryToken").Str("token", p.Token).Logger()

	recovery, err := s.recoveryRepository.GetByID(repository.DefaultNoTransaction, p.Token)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve recovery")
		return err
	}

	if !recoveryIsValid(recovery) {
		msg := "recovery token is expired"
		l.Info().Msg(msg)
		return coachee.MakeUnauthorized(errors.New(msg))
	}

	if len(p.Password) < 8 {
		msg := "password needs to be longer than 8 characters"
		l.Debug().Msg(msg)
		return coachee.MakeValidation(errors.New(msg))
	}

	hashedPass, err := auth.Hash(p.Password)
	if err != nil {
		l.Info().Err(err).Msg("failed to hash password")
		return coachee.MakeInternal(err)
	}

	customer, err := s.customerRepository.GetByID(repository.DefaultNoTransaction, recovery.CustomerID)
	if err != nil {
		l.Error().Err(err).Uint("CustomerId", recovery.CustomerID).Msg("failed to retrieve customer")
		return err
	}

	customer.Password = string(hashedPass)
	if err := s.customerRepository.Update(repository.DefaultNoTransaction, customer); err != nil {
		l.Error().Err(err).Uint("CustomerId", recovery.CustomerID).Msg("failed to update customer")
		return err
	}

	return nil
}

func recoveryIsValid(r *model.Recovery) bool {
	if r == nil {
		return false
	}
	return r.CreatedAt.Add(24 * time.Hour).After(time.Now())
}
