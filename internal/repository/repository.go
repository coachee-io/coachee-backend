package repository

import (
	"coachee-backend/internal/model"
)

// Coach is the repository to interact and persist coaches
type Coach interface {
	Begin() Transaction

	Create(tx Transaction, coach *model.Coach) error
	GetByID(tx Transaction, id uint) (*model.Coach, error)
	GetByEmail(tx Transaction, email string) (*model.Coach, error)
	GetByPage(tx Transaction, tag string, limit, page uint) ([]*model.Coach, error)
	Update(tx Transaction, coach *model.Coach) error
	Length(tx Transaction, tag string) (uint, error)
}

// Customer is the repository to interact and persist clients
type Customer interface {
	Begin() Transaction

	Create(tx Transaction, client *model.Customer) error
	GetByID(tx Transaction, id uint) (*model.Customer, error)
	GetByEmail(tx Transaction, email string) (*model.Customer, error)
	Update(tx Transaction, client *model.Customer) error
}

// Order is the repository to interact and persist orders
type Order interface {
	Begin() Transaction

	Create(transaction Transaction, client *model.Order) error
	GetByID(transaction Transaction, id uint) (*model.Order, error)
	Update(transaction Transaction, client *model.Order) error
}

// Recovery is the repository to interact and persist password recovery flows
type Recovery interface {
	Begin() Transaction

	Create(transaction Transaction, client *model.Recovery) error
	GetByID(transaction Transaction, id string) (*model.Recovery, error)
}

// CoachRecovery is the repository to interact and persist password recovery flows for a coach
type CoachRecovery interface {
	Begin() Transaction

	Create(transaction Transaction, client *model.CoachRecovery) error
	GetByID(transaction Transaction, id string) (*model.CoachRecovery, error)
}
