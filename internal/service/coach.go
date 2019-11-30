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

	coaches, err := s.coachRepository.GetByPage(repository.DefaultNoTransaction, p.Tag, limit, page)
	if err != nil {
		s.logger.Error().Err(err).Str("tags", p.Tag).Msg("failed to retrieve coaches")
		return nil, err
	}

	return CoachesToPayload(coaches), nil
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
	coach := &model.Coach{
		FirstName:   p.FirstName,
		LastName:    p.LastName,
		Email:       p.Email,
		Phone:       p.Phone,
		Tags:        p.Tags,
		Description: p.Description,
		City:        *p.City,
		Country:     *p.Country,
		Status:      model.StatusRegistered,
		Vat:         *p.Vat,
		IntroCall:   time.Unix(int64(p.IntroCall), 0),
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

	switch {
	case p.FirstName != nil:
		coach.FirstName = *p.FirstName
		fallthrough
	case p.LastName != nil:
		coach.LastName = *p.LastName
		fallthrough
	case p.Email != nil:
		coach.Email = *p.Email
		fallthrough
	case p.Phone != nil:
		coach.Phone = *p.Phone
		fallthrough
	case p.Tags != nil:
		coach.Tags = *p.Tags
		fallthrough
	case p.Description != nil:
		coach.Description = *p.Description
		fallthrough
	case p.City != nil:
		coach.City = *p.City
		fallthrough
	case p.Country != nil:
		coach.Country = *p.Country
		fallthrough
	case p.IntroCall != nil:
		coach.IntroCall = time.Unix(int64(*p.IntroCall), 0)
		fallthrough
	case p.StripeID != nil:
		coach.StripeID = *p.StripeID
		fallthrough
	case p.PictureURL != nil:
		coach.PictureUrl = *p.PictureURL
		fallthrough
	case p.Vat != nil:
		coach.Vat = *p.Vat
	}

	err := s.coachRepository.Update(repository.DefaultNoTransaction, coach)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to update coach")
		return err
	}

	return nil
}
