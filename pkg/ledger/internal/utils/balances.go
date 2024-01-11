package utils

import (
	"fmt"
	"strings"
)

// FetchAllBalances fetches all balances via the Ledger CLI.
func FetchAllBalances(filepath string) ([]string, error) {
	command := fmt.Sprintf("ledger balance -f %s", filepath)
	output, err := ExecuteShellCommand(command)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch balances: %w", err)
	}

	balances := strings.Split(output, "\n")
	return balances, nil
}
