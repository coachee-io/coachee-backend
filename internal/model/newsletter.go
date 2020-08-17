package model

import "time"

type Newsletter struct {
	ID    uint   `gorm:"primary_key"`
	Email string `gorm:"unique;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
