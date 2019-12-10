package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"errors"
)

// CreateCertification creates a certification for a coach
func (s *Service) CreateCertification(ctx context.Context, p *coachee.CreateCertificationPayload) (err error) {
	l := s.logger.With().Str("service", "CreateCertification").Logger()

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to retrieve coach")
		return err
	}

	c := model.Certification{}.New(p.Certification)
	coach.Certifications = append(coach.Certifications, c)

	err = s.coachRepository.Update(repository.DefaultNoTransaction, coach)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to update coach")
		return err
	}

	return nil
}

// deletes a certification for a coach
func (s *Service) DeleteCertification(ctx context.Context, p *coachee.DeleteCertificationPayload) (err error) {
	l := s.logger.With().Str("service", "DeleteCertification").Logger()

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to retrieve coach")
		return err
	}

	for i, c := range coach.Certifications {
		if c.ID == p.CertID {
			coach.Certifications = append(coach.Certifications[:i], coach.Certifications[i+1:]...)
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
