package utils_test

import (
	"htmx/pkg/ledger/internal/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccounts(t *testing.T) {
	t.Run("should return a slice of accounts", func(t *testing.T) {
		// given
		file := "./testdata/ledger2023.ledger"

		// when
		accounts, err := utils.Accounts(file)

		for _, account := range accounts {
			log.Printf(account.ToString())
		}

		// then
		assert.NoError(t, err)
		assert.Equal(t, 15, len(accounts))
	})
}
