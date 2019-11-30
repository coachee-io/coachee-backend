package service

import (
	"coachee-backend/gen/coachee"
	"context"
)

// creates a certification for a coach
func (s *Service) CreateCertification(ctx context.Context, p *coachee.CreateCertificationPayload) (err error) {
	s.logger.Print("coachee.CreateCertification")
	return
}

// deletes a certification for a coach
func (s *Service) DeleteCertification(ctx context.Context, p *coachee.DeleteCertificationPayload) (err error) {
	s.logger.Print("coachee.DeleteCertification")
	return
}
