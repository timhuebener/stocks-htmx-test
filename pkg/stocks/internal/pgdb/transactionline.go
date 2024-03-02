package pgdb

import (
	"context"
	"gorm.io/gorm"
	"htmx/pkg/otel/db/psql"
)

type TransactionLine struct {
	gorm.Model
	TransactionID uint
	AccountID     uint
	Debit         float64 // Amount debited from the account
	Credit        float64 // Amount credited to the account
}

// CreateTransactionLine - Create a new transactionLine
func CreateTransactionLine(ctx context.Context, transactionLine TransactionLine) (*TransactionLine, error) {
	return psql.Create[TransactionLine](ctx, transactionLine)
}

// GetTransactionLineByID - Get an transactionLine by ID
func GetTransactionLineByID(ctx context.Context, id uint) (*TransactionLine, error) {
	var transactionLine TransactionLine
	return psql.GetByID[TransactionLine, uint](ctx, transactionLine, id)
}

// GetTransactionLines - Get all transactionLines
func GetTransactionLines(ctx context.Context) ([]TransactionLine, error) {
	var transactionLines []TransactionLine
	return psql.Get[TransactionLine](ctx, transactionLines)
}

// UpdateTransactionLine - Update an existing transactionLine
func UpdateTransactionLine(ctx context.Context, transactionLine TransactionLine) (*TransactionLine, error) {
	return psql.Update[TransactionLine](ctx, transactionLine)
}

// DeleteTransactionLine - Delete a transactionLine
func DeleteTransactionLine(ctx context.Context, id uint) error {
	var transactionLine TransactionLine
	return psql.Delete[TransactionLine, uint](ctx, transactionLine, id)
}

func SeedTransactionLine(ctx context.Context) (*TransactionLine, error) {
	account, err := SeedAccount(ctx)
	if err != nil {
		return nil, err
	}
	transaction, err := SeedTransaction(ctx)
	if err != nil {
		return nil, err
	}
	a := TransactionLine{
		TransactionID: transaction.ID,
		AccountID:     account.ID,
		Debit:         100,
		Credit:        0,
	}
	return CreateTransactionLine(ctx, a)
}
