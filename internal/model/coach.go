package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
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
	ID          uint `gorm:"primary_key"`
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	StripeID    string
	Tags        string
	Description string `sql:"TYPE:text"`
	City        string
	Country     string
	PictureUrl  string
	Status      CoachStatus
	Vat         string
	IntroCall   time.Time

	Availability   Availabilities `sql:"TYPE:json"`
	Certifications Certifications `sql:"TYPE:json"`
	Programs       Programs       `sql:"TYPE:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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
