package pgdb

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"htmx/pkg/otel"
	"htmx/pkg/otel/log"
)

var DB *gorm.DB

func Connect(connectionString string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	return nil
}

func Migrate() error {
	return DB.AutoMigrate(&Account{}, &Transaction{}, &TransactionLine{})
}

func Seed() {
	a := Account{
		Name:    "ING",
		Type:    string(Asset),
		Balance: 0,
	}
	res, err := CreateAccount(&a)
	if err != nil {
		log.Fatal(context.TODO(), "unable to create account", otel.ErrorMsg.String(err.Error()))
	}
	log.Info(context.TODO(), fmt.Sprintf("%+v", res))
}
