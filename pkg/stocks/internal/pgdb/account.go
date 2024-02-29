package pgdb

import "gorm.io/gorm"

type Type string

const (
	Asset Type = "asset"
)

type Account struct {
	gorm.Model
	Name    string
	Type    string // assets, liabilities, equity, expenses, revenues
	Balance float64
}

// CreateAccount - Create a new account
func CreateAccount(account *Account) (*Account, error) {
	err := DB.Create(&account).Error
	return account, err
}

// GetAccountByID - Get a account by ID
func GetAccountByID(id uint) (*Account, error) {
	var account Account
	err := DB.First(&account, id).Error
	return &account, err
}

// GetTransactions - Get all transactions
func GetAccounts() ([]Account, error) {
	var account []Account
	res := DB.Find(&account)
	return account, res.Error
}

// UpdateAccount - Update an existing account
func UpdateAccount(account *Account) error {
	return DB.Save(&account).Error
}

// DeleteAccount - Delete a account
func DeleteAccount(id uint) error {
	var account Account
	return DB.Delete(&account, id).Error
}
