package service

import (
	"coachee-backend/gen/coachee"
	"context"
)

// creates a program for a coach
func (s *Service) CreateProgram(ctx context.Context, p *coachee.CreateProgramPayload) (err error) {
	s.logger.Print("coachee.CreateProgram")
	return
}

// deletes a program for a coach
func (s *Service) DeleteProgram(ctx context.Context, p *coachee.DeleteProgramPayload) (err error) {
	s.logger.Print("coachee.DeleteProgram")
	return
}
