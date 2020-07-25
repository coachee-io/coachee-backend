package model

import (
	"coachee-backend/gen/coachee"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/pborman/uuid"
)

// CoachStatus are the possible status for the coaches
type CoachStatus string

// types of coach status
const (
	StatusRegistered = CoachStatus("registered")
	StatusRejected   = CoachStatus("rejected")
	StatusInactive   = CoachStatus("inactive")
	StatusActive     = CoachStatus("active")
)

// Coach represents all information relative to the coach
type Coach struct {
	ID                uint `gorm:"primary_key"`
	FirstName         string
	LastName          string
	Email             string `gorm:"not null;unique"`
	Password          string
	Phone             string
	StripeID          string
	Tags              string
	Description       string `sql:"TYPE:text"`
	City              string
	Country           string
	PictureUrl        string
	Status            CoachStatus
	Vat               string
	IntroCall         time.Time
	FirstCallDuration int32
	VideoURL          string
	CardDescription   string

	Availability   Availabilities `sql:"TYPE:json"`
	Certifications Certifications `sql:"TYPE:json"`
	Programs       Programs       `sql:"TYPE:json"`

	TextAvailability   string `sql:"TYPE:text"`
	TextCertifications string `sql:"TYPE:text"`
	TextPrograms       string `sql:"TYPE:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Validate checks if the value is valid
func (s CoachStatus) Validate() error {
	// string casting is necessary to avoid validation recursion...
	if err := validation.Validate(string(s),
		validation.In(
			string(StatusRegistered),
			string(StatusRejected),
			string(StatusInactive),
			string(StatusActive),
		)); err != nil {
		return coachee.MakeValidation(err)
	}
	return nil
}

// Availabilities is an array of chunks of time
type Availabilities []*Availability

// Value overrides the default sql store func
func (i Availabilities) Value() (driver.Value, error) {
	b, err := json.Marshal(i)
	return string(b), err
}

// Scan overrides the default sql retrieval func
func (i *Availabilities) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &i)
}

// Certifications represents all certifications a coach has
type Certifications []*Certification

// Value overrides the default sql store func
func (i Certifications) Value() (driver.Value, error) {
	b, err := json.Marshal(i)
	return string(b), err
}

// Scan overrides the default sql retrieval func
func (i *Certifications) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &i)
}

// Programs represents all programs a coach has
type Programs []*Program

// Value overrides the default sql store func
func (i Programs) Value() (driver.Value, error) {
	b, err := json.Marshal(i)
	return string(b), err
}

// Scan overrides the default sql retrieval func
func (i *Programs) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &i)
}

// Availability represents a chunk of time when a coach is available
type Availability struct {
	ID    string
	Day   uint
	Start uint // in minutes
	End   uint
}

// Program represents a coach program
type Program struct {
	ID string

	Name             string
	NumberOfSessions uint
	Duration         uint // in minutes
	Description      string
	TotalPrice       uint // in cents
	TaxPercent       uint // per 10000
}

// Certification represents a coach certification
type Certification struct {
	ID string

	Title        string
	Description  string
	Institution  string
	DateAcquired time.Time
}

// New creates a new Availability from the payload
func (Availability) New(a *coachee.CreateAvailabilityPayload) *Availability {
	if a == nil {
		return nil
	}

	return &Availability{
		ID:    uuid.New(),
		Day:   a.WeekDay,
		Start: a.Start,
		End:   a.End,
	}
}

// New creates a new certification from the payload
func (Certification) New(c *coachee.Certification) *Certification {
	if c == nil {
		return nil
	}

	return &Certification{
		ID:           uuid.New(),
		Title:        c.Title,
		Description:  c.Description,
		Institution:  c.Institution,
		DateAcquired: time.Date(int(c.Year), time.Month(c.Month), 1, 0, 0, 0, 0, time.UTC),
	}
}

// New creates a new program from the payload
func (Program) New(p *coachee.Program) *Program {
	if p == nil {
		return nil
	}

	return &Program{
		ID:               uuid.New(),
		Name:             p.Name,
		NumberOfSessions: p.Sessions,
		Duration:         p.Duration,
		Description:      p.Description,
		TotalPrice:       p.TotalPrice,
		TaxPercent:       p.TaxPercent,
	}
}

// GetProgram
func (c Coach) GetProgram(id string) (*Program, error) {
	for _, program := range c.Programs {
		if program.ID == id {
			return program, nil
		}
	}
	return nil, coachee.MakeNotFound(errors.New("program not found"))
}
