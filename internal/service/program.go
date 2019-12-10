package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"errors"
)

// CreateProgram creates a program for a coach
func (s *Service) CreateProgram(ctx context.Context, p *coachee.CreateProgramPayload) error {
	l := s.logger.With().Str("service", "CreateProgram").Logger()

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to retrieve coach")
		return err
	}

	program := model.Program{}.New(p.Program)
	coach.Programs = append(coach.Programs, program)

	err = s.coachRepository.Update(repository.DefaultNoTransaction, coach)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to update coach")
		return err
	}

	return nil
}

// DeleteProgram deletes a program for a coach
func (s *Service) DeleteProgram(ctx context.Context, p *coachee.DeleteProgramPayload) error {
	l := s.logger.With().Str("service", "DeleteProgram").Logger()

	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		l.Error().Err(err).Uint("coachID", p.ID).Msg("failed to retrieve coach")
		return err
	}

	for i, program := range coach.Programs {
		if program.ID == p.ProgramID {
			coach.Programs = append(coach.Programs[:i], coach.Programs[i+1:]...)
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
