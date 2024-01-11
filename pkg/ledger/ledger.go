package ledger

import (
	"htmx/pkg/ledger/internal/utils"
	"htmx/pkg/ledger/models"
)

// Accounts fetches all accounts via the Ledger CLI.
// It takes a filepath as input and returns a slice of models.Account and an error.
func Accounts(filepath string) ([]models.Account, error) {
	return utils.Accounts(filepath)
}
