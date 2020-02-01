package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"context"
	"time"
)

// GetCoaches returns an array of coaches according to a tag and pagination
func (s *Service) GetCoaches(ctx context.Context, p *coachee.GetCoachesPayload) ([]*coachee.Coach, error) {
	var limit, page uint
	limit = 100 // default

	if p.Limit != nil {
		limit = *p.Limit
	}

	if p.Page != nil {
		page = *p.Page
	}
	var tag string
	if p.Tag != nil {
		tag = *p.Tag
	}

	coaches, err := s.coachRepository.GetByPage(repository.DefaultNoTransaction, tag, limit, page)
	if err != nil {
		s.logger.Error().Err(err).Str("tags", tag).Msg("failed to retrieve coaches")
		return nil, err
	}

	return CoachesToPayload(coaches), nil
}

// GetCoach returns a coach according to the id
func (s *Service) GetCoach(ctx context.Context, p *coachee.GetCoachPayload) (*coachee.Coach, error) {
	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		s.logger.Error().Err(err).Uint("id", p.ID).Msg("failed to retrieve coach")
		return nil, err
	}
	return CoachToPayload(coach), nil
}

// LenCoaches gives the number of coaches with a given tag
func (s *Service) LenCoaches(ctx context.Context, p *coachee.LenCoachesPayload) (uint, error) {
	count, err := s.coachRepository.Length(repository.DefaultNoTransaction, p.Tag)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get count of coaches")
		return 0, err
	}

	return count, err
}

// CreateCoaches creates a base coach
func (s *Service) CreateCoach(ctx context.Context, p *coachee.CreateCoachPayload) (uint, error) {
	var city, country, vat, textAvailability string
	if p.City != nil {
		city = *p.City
	}
	if p.Country != nil {
		country = *p.Country
	}
	if p.Vat != nil {
		vat = *p.Vat
	}
	if p.TextAvailability != nil {
		textAvailability = *p.TextAvailability
	}

	coach := &model.Coach{
		FirstName:          p.FirstName,
		LastName:           p.LastName,
		Email:              p.Email,
		Phone:              p.Phone,
		Tags:               p.Tags,
		Description:        p.Description,
		City:               city,
		Country:            country,
		Status:             model.StatusRegistered,
		Vat:                vat,
		IntroCall:          time.Unix(int64(p.IntroCall), 0),
		TextAvailability:   textAvailability,
		TextCertifications: p.TextCertifications,
		TextPrograms:       p.TextPrograms,
	}

	err := s.coachRepository.Create(repository.DefaultNoTransaction, coach)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create coach")
		return 0, err
	}

	return coach.ID, err
}

// UpdateCoaches updates a coach
func (s *Service) UpdateCoach(ctx context.Context, p *coachee.UpdateCoachPayload) error {
	coach := &model.Coach{
		ID: p.ID,
	}

	if p.FirstName != nil {
		coach.FirstName = *p.FirstName
	}
	if p.LastName != nil {
		coach.LastName = *p.LastName
	}
	if p.Email != nil {
		coach.Email = *p.Email
	}
	if p.Phone != nil {
		coach.Phone = *p.Phone
	}
	if p.Tags != nil {
		coach.Tags = *p.Tags
	}
	if p.Description != nil {
		coach.Description = *p.Description
	}
	if p.City != nil {
		coach.City = *p.City
	}
	if p.Country != nil {
		coach.Country = *p.Country
	}
	if p.IntroCall != nil {
		coach.IntroCall = time.Unix(int64(*p.IntroCall), 0)
	}
	if p.StripeID != nil {
		coach.StripeID = *p.StripeID
	}
	if p.PictureURL != nil {
		coach.PictureUrl = *p.PictureURL
	}
	if p.Vat != nil {
		coach.Vat = *p.Vat
	}

	err := s.coachRepository.Update(repository.DefaultNoTransaction, coach)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to update coach")
		return err
	}

	return nil
}

// RegisterStripeExpress registers a stripe express account and associates it to a coach
func (s *Service) RegisterStripeExpress(ctx context.Context, p *coachee.RegisterStripeExpressPayload) error {
	l := s.logger.With().Str("service", "RegisterStripeExpress").Logger()

	stripeID, err := s.stripe.RegisterStripeExpress(p.ExpressID)
	if err != nil {
		l.Error().Err(err).Msg("failed to register stripe express")
		return err
	}

	coachUpdate := &model.Coach{
		ID:       p.ID,
		StripeID: stripeID,
	}
	if err := s.coachRepository.Update(repository.DefaultNoTransaction, coachUpdate); err != nil {
		l.Error().Err(err).Str("stripeID", stripeID).Msg("failed to persist coach stripe id")
		return coachee.MakeInternal(err)
	}

	return nil
}
