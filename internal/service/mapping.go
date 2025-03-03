package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"fmt"
	"sort"
	"time"
)

func CoachesToPayload(coaches []*model.Coach) []*coachee.Coach {
	var res []*coachee.Coach
	for _, coach := range coaches {
		res = append(res, CoachToPayload(coach))
	}
	return res
}

func CoachToPayload(c *model.Coach) *coachee.Coach {
	if c == nil {
		return nil
	}

	return &coachee.Coach{
		ID:                c.ID,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		Tags:              c.Tags,
		Description:       c.Description,
		City:              c.City,
		Country:           c.Country,
		PictureURL:        c.PictureUrl,
		FirstCallDuration: c.FirstCallDuration,
		Certifications:    CertificationsToPayload(c.Certifications),
		Programs:          ProgramsToPayload(c.Programs),
		Availability:      AvailabilitiesToPayload(c.Availability),
		VideoURL:          c.VideoURL,
		CardDescription:   c.CardDescription,
	}
}

func FullCoachToPayload(c *model.Coach) *coachee.FullCoach {
	if c == nil {
		return nil
	}

	return &coachee.FullCoach{
		ID:                c.ID,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		Email:             c.Email,
		Phone:             c.Phone,
		StripeID:          c.StripeID,
		Tags:              c.Tags,
		Description:       c.Description,
		City:              c.City,
		Country:           c.Country,
		PictureURL:        c.PictureUrl,
		Status:            string(c.Status),
		Vat:               c.Vat,
		IntroCall:         int(c.IntroCall.Unix()),
		FirstCallDuration: c.FirstCallDuration,
		VideoURL:          c.VideoURL,
		CardDescription:   c.CardDescription,
		Availability:      AvailabilitiesToPayload(c.Availability),
		Certifications:    CertificationsToPayload(c.Certifications),
		Programs:          ProgramsToPayload(c.Programs),
	}
}

func CertificationsToPayload(c model.Certifications) []*coachee.Certification {
	if c == nil {
		return nil
	}

	var certs []*coachee.Certification
	for _, certification := range c {
		certs = append(certs, CertificationToPayload(certification))
	}

	return certs
}

func CertificationToPayload(c *model.Certification) *coachee.Certification {
	if c == nil {
		return nil
	}

	return &coachee.Certification{
		ID:          &c.ID,
		Title:       c.Title,
		Description: c.Description,
		Institution: c.Institution,
		Month:       uint(c.DateAcquired.Month()),
		Year:        uint(c.DateAcquired.Year()),
	}
}

func ProgramsToPayload(p model.Programs) []*coachee.Program {
	if p == nil {
		return nil
	}

	var programs []*coachee.Program
	for _, program := range p {
		programs = append(programs, ProgramToPayload(program))
	}

	return programs
}

func ProgramToPayload(p *model.Program) *coachee.Program {
	if p == nil {
		return nil
	}

	var extra []*coachee.Session
	for _, s := range p.ExtraSessions {
		extra = append(extra, &coachee.Session{
			Sessions: s.NumberOfSessions,
			Duration: s.Duration,
		})
	}

	return &coachee.Program{
		ID:            &p.ID,
		Name:          p.Name,
		Sessions:      p.NumberOfSessions,
		Duration:      p.Duration,
		Description:   p.Description,
		TotalPrice:    p.TotalPrice,
		TaxPercent:    p.TaxPercent,
		ExtraSessions: extra,
	}
}

func AvailabilitiesToPayload(a model.Availabilities) []*coachee.Availability {
	if a == nil {
		return nil
	}

	var availabilities []*coachee.Availability
	for _, availability := range a {
		availabilities = append(availabilities, AvailabilityToPayload(availability))
	}
	sort.Slice(availabilities, func(i, j int) bool {
		if availabilities[i].WeekDay == availabilities[j].WeekDay {
			return availabilities[i].Start < availabilities[j].End
		}
		return availabilities[i].WeekDay < availabilities[j].WeekDay
	})

	return availabilities
}

func AvailabilityToPayload(a *model.Availability) *coachee.Availability {
	if a == nil {
		return nil
	}

	startHour := a.Start / 60
	startMinutes := a.Start % 60
	endHour := a.End / 60
	endMinutes := a.End % 60
	weekDay := time.Weekday(a.Day).String()[:3]
	dateText := fmt.Sprintf("%s %d:%02d-%d:%02d", weekDay, startHour, startMinutes, endHour, endMinutes)

	return &coachee.Availability{
		ID:        a.ID,
		WeekDay:   a.Day,
		Start:     float64(a.Start) / 60,
		End:       float64(a.End) / 60,
		DateLabel: dateText,
	}
}
