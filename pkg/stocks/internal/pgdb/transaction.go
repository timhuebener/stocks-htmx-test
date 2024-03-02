package pgdb

import (
	"context"
	"gorm.io/gorm"
	"htmx/pkg/otel/db/psql"
	"time"
)

type Transaction struct {
	gorm.Model
	Description string
	Date        time.Time
	Lines       []TransactionLine
}

// CreateTransaction - Create a new transaction
func CreateTransaction(ctx context.Context, transaction Transaction) (*Transaction, error) {
	return psql.Create[Transaction](ctx, transaction)
}

// GetTransactionByID - Get an transaction by ID
func GetTransactionByID(ctx context.Context, id uint) (*Transaction, error) {
	var transaction Transaction
	return psql.GetByID[Transaction, uint](ctx, transaction, id)
}

// GetTransactions - Get all transactions
func GetTransactions(ctx context.Context) ([]Transaction, error) {
	var transactions []Transaction
	return psql.Get[Transaction](ctx, transactions)
}

// UpdateTransaction - Update an existing transaction
func UpdateTransaction(ctx context.Context, transaction Transaction) (*Transaction, error) {
	return psql.Update[Transaction](ctx, transaction)
}

// DeleteTransaction - Delete a transaction
func DeleteTransaction(ctx context.Context, id uint) error {
	var transaction Transaction
	return psql.Delete[Transaction, uint](ctx, transaction, id)
}

func SeedTransaction(ctx context.Context) (*Transaction, error) {
	a := Transaction{
		Description: "Test Transaction",
		Date:        time.Now(),
	}
	return CreateTransaction(ctx, a)
}
