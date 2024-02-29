package pgdb

import (
	"gorm.io/gorm"
)

type TransactionLine struct {
	gorm.Model
	TransactionID uint
	AccountID     uint
	Debit         float64 // Amount debited from the account
	Credit        float64 // Amount credited to the account
}

// CreateTransactionLine - Create a new transactionLine
func CreateTransactionLine(transactionLine *TransactionLine) (*TransactionLine, error) {
	err := DB.Create(&transactionLine).Error
	return transactionLine, err
}

// GetTransactionLines - Get all transactionLines
func GetTransactionLines() ([]TransactionLine, error) {
	var transactionLines []TransactionLine
	res := DB.Find(&transactionLines)
	return transactionLines, res.Error
}

// GetTransactionLineByID - Get a transactionLine by ID
func GetTransactionLineByID(id uint) (*TransactionLine, error) {
	var transactionLine TransactionLine
	err := DB.First(&transactionLine, id).Error
	return &transactionLine, err
}
