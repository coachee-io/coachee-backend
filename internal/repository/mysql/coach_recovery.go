package mysql

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"

	"github.com/jinzhu/gorm"
)

// CoachRecoveryRepository is the repository to access and persist clientes
type CoachRecoveryRepository struct {
	db *gorm.DB
}

// NewCoachRecoveryRepository initializes a CoachRecoveryRepository
func NewCoachRecoveryRepository(db *gorm.DB) *CoachRecoveryRepository {
	return &CoachRecoveryRepository{db: db}
}

// Begin starts a new Transaction
func (r CoachRecoveryRepository) Begin() repository.Transaction {
	return newTransaction(r.db)
}

// Create persists a client
func (r CoachRecoveryRepository) Create(transaction repository.Transaction, client *model.CoachRecovery) error {
	tx := r.checkTransaction(transaction)

	// create recovery
	if err := tx.Create(client).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// GetByID returns a client by id
func (r CoachRecoveryRepository) GetByID(transaction repository.Transaction, id string) (*model.CoachRecovery, error) {
	tx := r.checkTransaction(transaction)

	var client model.CoachRecovery
	if err := tx.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, parseError(err)
	}

	return &client, nil
}

// checkTransaction returns either a gorm.DB pointer either as a transaction or as a direct db access
// this depends on the if the repository.Transaction is a transaction or a NoTransaction
func (r CoachRecoveryRepository) checkTransaction(in repository.Transaction) *gorm.DB {
	tx, ok := in.(*transaction)
	if !ok {
		return r.db
	}
	return tx.db
}
