package mysql

import (
	"coachee-backend/internal/model"
	"coachee-backend/internal/repository"

	"github.com/jinzhu/gorm"
)

// NewsletterRepository is the repository to newsletter emails
type NewsletterRepository struct {
	db *gorm.DB
}

// NewNewsletterRepository initializes a NewsletterRepository
func NewNewsletterRepository(db *gorm.DB) *NewsletterRepository {
	return &NewsletterRepository{db: db}
}

// Begin starts a new Transaction
func (r NewsletterRepository) Begin() repository.Transaction {
	return newTransaction(r.db)
}

// Create persists a client
func (r NewsletterRepository) Create(transaction repository.Transaction, client *model.Newsletter) error {
	tx := r.checkTransaction(transaction)

	// create recovery
	if err := tx.Create(client).Error; err != nil {
		return parseError(err)
	}

	return nil
}

// checkTransaction returns either a gorm.DB pointer either as a transaction or as a direct db access
// this depends on the if the repository.Transaction is a transaction or a NoTransaction
func (r NewsletterRepository) checkTransaction(in repository.Transaction) *gorm.DB {
	tx, ok := in.(*transaction)
	if !ok {
		return r.db
	}
	return tx.db
}
