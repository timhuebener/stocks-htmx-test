package psql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(connectionString string) error {
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	return nil
}

func Migrate(dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}

// Create - Create a new entry
func Create[T interface{}](value T) (*T, error) {
	err := db.Create(&value).Error
	return &value, err
}

// Get - Get all entries
func Get[T interface{}](dest []T) ([]T, error) {
	res := db.Find(&dest)
	return dest, res.Error
}

// GetByID - Get an entry by ID
func GetByID[T interface{}, C interface{}](dest T, conds ...C) (*T, error) {
	err := db.First(&dest, conds).Error
	return &dest, err
}

// Update - Update an existing entry
func Update[T interface{}](value T) (*T, error) {
	err := db.Save(&value).Error
	return &value, err
}

// Delete - Delete an entry
func Delete[T interface{}, C interface{}](value T, conds ...C) error {
	return db.Delete(&value, conds).Error
}
