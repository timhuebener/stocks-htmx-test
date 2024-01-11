package utils

import (
	"fmt"
	"htmx/pkg/ledger/models"
	"strings"
)

type AccountConnection struct {
	Name   string
	Parent string
	Child  string
}

func Accounts(filepath string) ([]models.Account, error) {
	accountStrings, err := read(filepath)
	if err != nil {
		return nil, err
	}
	accounts := parse(accountStrings)
	return accounts, nil
}

func read(filepath string) ([]string, error) {
	command := fmt.Sprintf("ledger accounts -f %s", filepath)
	output, err := ExecuteShellCommand(command)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch accounts: %w", err)
	}
	accounts := strings.Split(output, "\n")
	return accounts, nil
}

func parse(input []string) []models.Account {
	var accounts []models.Account
	for _, str := range input {
		accountParts := strings.Split(str, ":")
		for i := range accountParts {
			if accountParts[i] == "" {
				continue
			}
			account := findAccount(accountParts[i], accounts)
			if account == nil {
				account = &models.Account{
					Name: accountParts[i],
				}
			} else {
				account.SubAccounts = append(account.SubAccounts, models.Account{Name: accountParts[i]})
			}
			accounts = append(accounts, *account)
		}
	}
	return accounts
}

func findAccount(name string, accounts []models.Account) *models.Account {
	for i := range accounts {
		if accounts[i].Name == name {
			return &accounts[i]
		}
		for i := range accounts[i].SubAccounts {
			if accounts[i].SubAccounts[i].Name == name {
				return &accounts[i].SubAccounts[i]
			}
		}
	}
	return nil
}
