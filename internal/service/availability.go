package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"errors"
)

// CreateAvailability creates an availability for a coach
func (s *Service) CreateAvailability(ctx context.Context, p *coachee.CreateAvailabilityPayload) error {
	l := s.logger.With().Str("service", "CreateAvailability").Logger()

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to retrieve coach")
		return err
	}

	a := model.Availability{}.New(p.Availability)
	coach.Availability = append(coach.Availability, a)

	err = s.coachRepository.Update(repository.DefaultNoTransaction, coach)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to update coach")
		return err
	}

	return nil
}

// DeleteAvailability deletes an availability for a coach
func (s *Service) DeleteAvailability(ctx context.Context, p *coachee.DeleteAvailabilityPayload) error {
	l := s.logger.With().Str("service", "DeleteAvailability").Logger()

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to retrieve coach")
		return err
	}

	for i, a := range coach.Availability {
		if a.ID == p.AvID {
			coach.Availability = append(coach.Availability[:i], coach.Availability[i+1:]...)
			err = s.coachRepository.Update(repository.DefaultNoTransaction, coach)
			if err != nil {
				l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to update coach")
				return err
			}
			return nil
		}
	}

	l.Info().Uint("coachID", p.ID).Msg("coach not found")
	return coachee.MakeNotFound(errors.New("coach not found"))
}
