package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
)

// RegisterNewsletterEmail registers an email to our newsletter
func (s *Service) RegisterNewsletterEmail(ctx context.Context, p *coachee.RegisterNewsletterEmailPayload) (err error) {
	return s.newsletterRepository.Create(repository.DefaultNoTransaction, &model.Newsletter{Email: p.Email})
}
