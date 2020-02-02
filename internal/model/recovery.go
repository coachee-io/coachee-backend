package model

import (
	"time"

	"github.com/pborman/uuid"
)

// Recovery is the struct used for the password recovery flow
type Recovery struct {
	ID         string `gorm:"primary_key"`
	CustomerID uint
	CreatedAt  time.Time
}

func (r *Recovery) BeforeCreate() {
	r.ID = uuid.New()
}

// CoachRecovery is the struct used for the password recovery flow of a coach
type CoachRecovery struct {
	ID        string `gorm:"primary_key"`
	CoachID   uint
	CreatedAt time.Time
}

func (r *CoachRecovery) BeforeCreate() {
	r.ID = uuid.New()
}
