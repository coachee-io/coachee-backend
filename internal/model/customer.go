package model

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Customer struct {
	ID        uint `gorm:"primary_key"`
	StripeID  string
	FirstName string
	LastName  string
	Email     string `gorm:"not null;unique"`
	Password  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (c *Customer) Validate() error {
	if c == nil {
		return errors.New("nil client")
	}
	return validation.ValidateStruct(c,
		validation.Field(&c.Email, is.Email, validation.Required),
		validation.Field(&c.FirstName, validation.Required),
		validation.Field(&c.LastName, validation.Required),
		validation.Field(&c.Password, validation.Required))
}
