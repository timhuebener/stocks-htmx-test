package models

import "fmt"

type Account struct {
	Name        string
	SubAccounts []Account
}

func (a *Account) ToString() string {
	subAccounts := ""
	for _, sa := range a.SubAccounts {
		subAccounts += sa.Name + ", "
	}
	return fmt.Sprintf("{ %s [ %s ]}", a.Name, subAccounts)
}
