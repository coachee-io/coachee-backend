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
