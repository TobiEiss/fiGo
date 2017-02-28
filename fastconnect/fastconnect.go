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

type Task struct {
	AccountID            string `json:"account_id,omitempty"`
	IsEnded              bool   `json:"is_ended,omitempty"`
	IsErrpneous          bool   `json:"is_erroneous,omitempty"`
	IsWaitingForPin      bool   `json:"is_waiting_for_pin,omitempty"`
	IsWaitingForResponse bool   `json:"is_waiting_for_response,omitempty"`
	Message              string `json:"message,omitempty"`
}
