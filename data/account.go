package data

import "errors"

// Account contains the access credentials required to get a temporary session using MFA, along with a name to easily identify it with
type Account struct {
	Name      string
	AccessKey string
	SecretKey string
	MFASerial string

	Roles []AssumableRole
}

// GetRoleByName gets a role from the list with the given name, or returns a not found error
func (sd *Account) GetRoleByName(name string) (role AssumableRole, err error) {
	return role, errors.New("METHOD NOT IMPLEMENTED")
}
