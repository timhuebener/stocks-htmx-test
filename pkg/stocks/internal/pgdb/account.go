package pgdb

import (
	"context"
	"gorm.io/gorm"
	"htmx/pkg/otel/db/psql"
)

type Type string

const (
	Asset Type = "asset"
)

type Account struct {
	gorm.Model
	Name string
	Type string // assets, liabilities, equity, expenses, revenues
}

// CreateAccount - Create a new account
func CreateAccount(ctx context.Context, account Account) (*Account, error) {
	return psql.Create[Account](ctx, account)
}

// GetAccountByID - Get an account by ID
func GetAccountByID(ctx context.Context, id uint) (*Account, error) {
	var account Account
	return psql.GetByID[Account, uint](ctx, account, id)
}

// GetTransactions - Get all transactions
func GetAccounts(ctx context.Context) ([]Account, error) {
	var accounts []Account
	return psql.Get[Account](ctx, accounts)
}

// UpdateAccount - Update an existing account
func UpdateAccount(ctx context.Context, account Account) (*Account, error) {
	return psql.Update[Account](ctx, account)
}

// DeleteAccount - Delete a account
func DeleteAccount(ctx context.Context, id uint) error {
	var account Account
	return psql.Delete[Account, uint](ctx, account, id)
}

func SeedAccount(ctx context.Context) (*Account, error) {
	a := Account{
		Name: "Test Account",
		Type: string(Asset),
	}
	return CreateAccount(ctx, a)
}
