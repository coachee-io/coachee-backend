package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"
	"coachee-backend/pkg/slack"
	"context"
	"errors"
	"fmt"
	"time"
)

const firstCallDuration = 30

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
	var showAll bool
	if p.ShowAll != nil {
		showAll = *p.ShowAll
	}

	coaches, err := s.coachRepository.GetByPage(repository.DefaultNoTransaction, tag, limit, page, showAll)
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

// AdminGetCoach returns a coach according to the id. it will return most coach info
func (s *Service) AdminGetCoach(ctx context.Context, p *coachee.AdminGetCoachPayload) (*coachee.FullCoach, error) {
	coach, err := s.coachRepository.GetByID(repository.DefaultNoTransaction, p.ID)
	if err != nil {
		s.logger.Error().Err(err).Uint("id", p.ID).Msg("failed to retrieve coach")
		return nil, err
	}
	return FullCoachToPayload(coach), nil
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
	l := s.logger.With().Str("service", "CreateCoach").Logger()
	if !p.AcceptTerms {
		msg := "In order to proceed, please read and accept terms and conditions"
		l.Info().Msg(msg)
		return 0, coachee.MakeValidation(errors.New(msg))
	}

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

	if len(p.Password) < 8 {
		msg := "password needs to be longer than 8 characters"
		l.Debug().Msg(msg)
		return 0, coachee.MakeValidation(errors.New(msg))
	}

	hashedPass, err := auth.Hash(p.Password)
	if err != nil {
		l.Info().Err(err).Msg("failed to hash password")
		return 0, coachee.MakeValidation(err)
	}

	coach := &model.Coach{
		FirstName:          p.FirstName,
		LastName:           p.LastName,
		Email:              p.Email,
		Password:           string(hashedPass),
		Phone:              p.Phone,
		Tags:               p.Tags,
		Description:        p.Description,
		City:               city,
		Country:            country,
		Status:             model.StatusRegistered,
		Vat:                vat,
		IntroCall:          time.Unix(int64(p.IntroCall), 0),
		FirstCallDuration:  firstCallDuration,
		TextAvailability:   textAvailability,
		TextCertifications: p.TextCertifications,
		TextPrograms:       p.TextPrograms,
	}

	err = s.coachRepository.Create(repository.DefaultNoTransaction, coach)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create coach")
		return 0, err
	}

	msg := fmt.Sprintf("Coach %d %s %s has signed up.", coach.ID, coach.FirstName, coach.LastName)
	if err := s.slack.Post(slack.SimpleMessage(msg)); err != nil {
		l.Error().Err(err).Msg("failed to send slack message")
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
	if p.FirstCallDuration != nil {
		coach.FirstCallDuration = *p.FirstCallDuration
	}
	if p.Status != nil {
		coach.Status = model.CoachStatus(*p.Status)
		if err := coach.Status.Validate(); err != nil {
			return err
		}
	}
	if p.VideoURL != nil {
		coach.VideoURL = *p.VideoURL
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
	l := s.logger.With().Str("service", "RegisterStripeExpress").Str("code", p.AuthorizationCode).Logger()
	l.Debug().Msg("endpoint called")

	stripeID, err := s.stripe.RegisterStripeExpress(p.AuthorizationCode)
	if err != nil {
		l.Error().Err(err).Msg("failed to register stripe express")
		return err
	}

	coachUpdate := &model.Coach{
		ID:       p.ID,
		Status:   model.StatusActive,
		StripeID: stripeID,
	}
	if err := s.coachRepository.Update(repository.DefaultNoTransaction, coachUpdate); err != nil {
		l.Error().Err(err).Str("stripeID", stripeID).Msg("failed to persist coach stripe id")
		return coachee.MakeInternal(err)
	}

	msg := fmt.Sprintf("Coach %d has finalized the registration.", p.ID)
	if err := s.slack.Post(slack.SimpleMessage(msg)); err != nil {
		l.Error().Err(err).Msg("failed to send stripe message")
	}

	return nil
}

// LoginCoach returns a stripe express login link
func (s *Service) LoginCoach(ctx context.Context, p *coachee.LoginCoachPayload) (res *coachee.LoginCoachResult, err error) {
	l := s.logger.With().Str("service", "LoginCoach").Str("email", p.Email).Logger()

	coach, err := s.coachRepository.GetByEmail(repository.DefaultNoTransaction, p.Email)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve coach")
		return nil, coachee.MakeValidation(errors.New("your email or password is wrong"))
	}

	err = auth.VerifyPassword(coach.Password, p.Password)
	if err != nil {
		l.Debug().Err(err).Msg("failed to authenticate")
		return nil, coachee.MakeValidation(errors.New("your email or password is wrong"))
	}

	if coach.StripeID == "" {
		l.Info().Msg("empty stripe expressID")
		return nil, coachee.MakeUnauthorized(errors.New("stripe express account has not been set up yet"))
	}

	url, err := s.stripe.LoginStripeExpress(coach.StripeID)
	if err != nil {
		l.Error().Err(err).Msg("failed to retrieve stripe express account login url")
		return nil, coachee.MakeInternal(err)
	}

	return &coachee.LoginCoachResult{URL: url}, nil
}
