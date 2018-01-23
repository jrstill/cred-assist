package data

import "errors"

// StoredData contains the list of accounts, along with fields for reselecting recent items, that the app will use internally as well as to persist data
type StoredData struct {
	Accounts                []Account
	LastSelectedAccountName string
	LastSelectedRoleName    string
}

// GetAccountByName gets an account from the list with the given name, or returns a not found error
func (sd *StoredData) GetAccountByName(name string) (account Account, err error) {
	return account, errors.New("METHOD NOT IMPLEMENTED")
}
