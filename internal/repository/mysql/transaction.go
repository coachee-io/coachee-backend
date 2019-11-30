package mysql

import "github.com/jinzhu/gorm"

type transaction struct {
	db *gorm.DB
}

// NewTransaction initializes a new Transaction
func newTransaction(db *gorm.DB) *transaction {
	return &transaction{db: db.Begin()}
}

// Commit completes the transaction
func (tx transaction) Commit() error {
	return tx.db.Commit().Error
}

// Rollback cancels the transaction
func (tx transaction) Rollback() error {
	return tx.db.Rollback().Error
}
