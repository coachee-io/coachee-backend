package repository

import (
	"coachee-backend/internal/model"
)

// Coach is the repository to interact and persist coaches
type Coach interface {
	Begin() Transaction

	Create(tx Transaction, coach *model.Coach) error
	GetByID(tx Transaction, id uint) (*model.Coach, error)
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
