package data

// Account contains the access credentials required to get a temporary session using MFA, along with a name to easily identify it with
type Account struct {
	Name      string
	AccessKey string
	SecretKey string
	MFASerial string

	Roles []AssumableRole
}
