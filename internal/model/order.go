package model

import "time"

// Order defines a customers intent on buying a coach program
type Order struct {
	ID uint `gorm:"primary_key"`

	CoachID         uint
	CustomerID      uint
	ProgramID       string
	PaymentIntentID string

	Amount       uint
	TaxPercent   uint
	Status       string
	IntroCall    time.Time
	Observations string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
