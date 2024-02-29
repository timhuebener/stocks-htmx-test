package pgdb

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Description string
	Date        time.Time
	Lines       []TransactionLine
}

// CreateTransaction - Create a new transaction
func CreateTransaction(transaction *Transaction) (*Transaction, error) {
	err := DB.Create(&transaction).Error
	return transaction, err
}

// GetTransactions - Get all transactions
func GetTransactions() ([]Transaction, error) {
	var transactions []Transaction
	res := DB.Find(&transactions)
	return transactions, res.Error
}

// GetTransactionByID - Get a transaction by ID
func GetTransactionByID(id uint) (*Transaction, error) {
	var transaction Transaction
	err := DB.First(&transaction, id).Error
	return &transaction, err
}
