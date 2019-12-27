package mysql

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"

	"github.com/jinzhu/gorm"
)

// CustomerRepository is the repository to access and persist clientes
type CustomerRepository struct {
	db *gorm.DB
}

// NewClientRepository initializes a CustomerRepository
func NewClientRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// Begin starts a new Transaction
func (r CustomerRepository) Begin() repository.Transaction {
	return newTransaction(r.db)
}

// Create persists a client
func (r CustomerRepository) Create(transaction repository.Transaction, client *model.Customer) error {
	tx := r.checkTransaction(transaction)

	// create customer
	if err := tx.Create(client).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// GetByID returns a client by id
func (r CustomerRepository) GetByID(transaction repository.Transaction, id uint) (*model.Customer, error) {
	tx := r.checkTransaction(transaction)

	var client model.Customer
	if err := tx.First(&client, id).Error; err != nil {
		return nil, parseError(err)
	}

	return &client, nil
}

// GetByEmail returns a client by id
func (r CustomerRepository) GetByEmail(transaction repository.Transaction, email string) (*model.Customer, error) {
	tx := r.checkTransaction(transaction)

	var client model.Customer
	if err := tx.Where("email = ?", email).First(&client).Error; err != nil {
		return nil, parseError(err)
	}

	return &client, nil
}

// Update updates the changed fields in the db
func (r CustomerRepository) Update(transaction repository.Transaction, client *model.Customer) error {
	tx := r.checkTransaction(transaction)

	if err := tx.Model(client).Update(client).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// checkTransaction returns either a gorm.DB pointer either as a transaction or as a direct db access
// this depends on the if the repository.Transaction is a transaction or a NoTransaction
func (r CustomerRepository) checkTransaction(in repository.Transaction) *gorm.DB {
	tx, ok := in.(*transaction)
	if !ok {
		return r.db
	}
	return tx.db
}
