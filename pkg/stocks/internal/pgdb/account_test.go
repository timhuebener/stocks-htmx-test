package pgdb_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"htmx/pkg/db/psql"
	"htmx/pkg/stocks/internal/pgdb"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	conn := fmt.Sprintf("host=localhost user=%s dbname=%s sslmode=disable password=%s", "myuser", "mydatabase", "mypassword")
	if err := psql.Connect(conn); err != nil {
		t.Fatal("Failed to connect to database:", err)
	}

	t.Run("Create", func(t *testing.T) {
		// given
		a := pgdb.Account{
			Name:    "Create",
			Type:    "Test",
			Balance: 69420,
		}
		// when
		res, err := pgdb.CreateAccount(context.TODO(), a)
		// then
		assert.NoError(t, err)
		assert.Equal(t, a.Name, res.Name)
	})
	t.Run("GetById", func(t *testing.T) {
		// given
		a := pgdb.Account{
			Name:    "GetById",
			Type:    "Test",
			Balance: 69420,
		}
		res, err := pgdb.CreateAccount(context.TODO(), a)
		// when
		acc, err := pgdb.GetAccountByID(context.TODO(), res.ID)
		// then
		assert.NoError(t, err)
		assert.Equal(t, a.Name, acc.Name)
	})
	t.Run("Get", func(t *testing.T) {
		// given
		a := pgdb.Account{
			Name:    "Get",
			Type:    "Test",
			Balance: 69420,
		}
		_, err := pgdb.CreateAccount(context.TODO(), a)
		_, err = pgdb.CreateAccount(context.TODO(), a)
		// when
		acc, err := pgdb.GetAccounts(context.TODO())
		// then
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(acc), 2)
	})
	t.Run("Update", func(t *testing.T) {
		// given
		a := pgdb.Account{
			Name:    "Update",
			Type:    "Test",
			Balance: 69420,
		}
		res, err := pgdb.CreateAccount(context.TODO(), a)
		// when
		res.Name = "Updated"
		acc, err := pgdb.UpdateAccount(context.TODO(), *res)
		// then
		assert.NoError(t, err)
		assert.Equal(t, "Updated", acc.Name)
	})
	t.Run("Delete", func(t *testing.T) {
		// given
		a := pgdb.Account{
			Name:    "Delete",
			Type:    "Test",
			Balance: 69420,
		}
		res, err := pgdb.CreateAccount(context.TODO(), a)
		// when
		err = pgdb.DeleteAccount(context.TODO(), res.ID)
		// then
		assert.NoError(t, err)
	})
	acc, err := pgdb.GetAccounts(context.TODO())
	if err != nil {
		t.Fatal("unable to fetch all accounts")
	}
	for _, a := range acc {
		err = pgdb.DeleteAccount(context.TODO(), a.ID)
		if err != nil {
			t.Fatal("unable to delete account")
		}
	}
}
