package model

import "time"

// OrderStatus is an order status
type OrderStatus string

// Order status
const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusCaptured  OrderStatus = "captured"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusDeclined  OrderStatus = "declined"
)

// Order defines a customers intent on buying a coach program
type Order struct {
	ID uint `gorm:"primary_key"`

	CoachID         uint
	CustomerID      uint
	ProgramID       string
	PaymentIntentID string

	Amount       uint
	TaxPercent   uint
	Status       OrderStatus
	IntroCall    time.Time
	Observations string `sql:"TYPE:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
