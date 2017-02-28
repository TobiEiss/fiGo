package fastconnect

import "errors"

var (
	// ErrRecoveryPassword cant find a "recovery_password" in JSON
	ErrRecoveryPassword = errors.New("figo-fastconnect: extract recovery password failed")
)

// FigoUser represent a user
type FigoUser struct {
	Email    string
	Username string
	Password string
}

// BankCredentials represent the BankCredentials
type BankCredentials struct {
	BankCode    string
	Country     string
	Credentials []string
}
