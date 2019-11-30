package service

import (
	"coachee-backend/gen/coachee"
	"context"
)

// creates an availability for a coach
func (s *Service) CreateAvailability(ctx context.Context, p *coachee.CreateAvailabilityPayload) (err error) {
	s.logger.Print("coachee.CreateAvailability")
	return
}

// deletes an availability for a coach
func (s *Service) DeleteAvailability(ctx context.Context, p *coachee.DeleteAvailabilityPayload) (err error) {
	s.logger.Print("coachee.DeleteAvailability")
	return
}
