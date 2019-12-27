package mysql

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"

	"github.com/jinzhu/gorm"
)

// OrderRepository is the repository to access and persist clientes
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository initializes a OrderRepository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Begin starts a new Transaction
func (r OrderRepository) Begin() repository.Transaction {
	return newTransaction(r.db)
}

// Create persists a client
func (r OrderRepository) Create(transaction repository.Transaction, client *model.Order) error {
	tx := r.checkTransaction(transaction)

	// create order
	if err := tx.Create(client).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// GetByID returns a client by id
func (r OrderRepository) GetByID(transaction repository.Transaction, id uint) (*model.Order, error) {
	tx := r.checkTransaction(transaction)

	var client model.Order
	if err := tx.First(&client, id).Error; err != nil {
		return nil, parseError(err)
	}

	return &client, nil
}

// Update updates the changed fields in the db
func (r OrderRepository) Update(transaction repository.Transaction, client *model.Order) error {
	tx := r.checkTransaction(transaction)

	if err := tx.Model(client).Update(client).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// checkTransaction returns either a gorm.DB pointer either as a transaction or as a direct db access
// this depends on the if the repository.Transaction is a transaction or a NoTransaction
func (r OrderRepository) checkTransaction(in repository.Transaction) *gorm.DB {
	tx, ok := in.(*transaction)
	if !ok {
		return r.db
	}
	return tx.db
}
