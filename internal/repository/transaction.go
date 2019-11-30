package repository

import "fmt"

// Transaction abstracts a db transaction to be managed by the service layer
//go:generate mockgen -destination ../../mocks/repository_transaction_mock.go -package mocks git.perkbox.io/backend-services/billing-invoices/internal/repository Transaction
type Transaction interface {
	Commit() error
	Rollback() error
}

// NoTransaction is the structure to be used when using the repositories without Transactions
type NoTransaction struct{}

// DefaultNoTransaction is the default variable to be used when using the repositories without Transactions
var DefaultNoTransaction = NoTransaction{}

// Commit exist only to implement the Transaction interface. DO NOT USE.
func (n NoTransaction) Commit() error {
	return nil
}

// Rollback exist only to implement the Transaction interface. DO NOT USE.
func (n NoTransaction) Rollback() error {
	return fmt.Errorf("cannot Rollback with NoTransaction: the transaction is already committed")
}
