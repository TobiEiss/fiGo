package fastconnect

import "errors"

var (
	// ErrRecoveryPassword cant find a "recovery_password" in JSON
	ErrRecoveryPassword = errors.New("figo-fastconnect: extract recovery password failed")
)

// FigoUser represent a user
type FigoUser struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Username string `json:"username,omitempty" bson:"emusernameail,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

// BankCredentials represent the BankCredentials
type BankCredentials struct {
	BankCode    string   `json:"bankCode,omitempty" bson:"bankCode,omitempty"`
	Country     string   `json:"country,omitempty" bson:"country,omitempty"`
	Credentials []string `json:"credentials,omitempty" bson:"credentials,omitempty"`
}

type Task struct {
	AccountID            string `json:"account_id,omitempty"`
	IsEnded              bool   `json:"is_ended,omitempty"`
	IsErrpneous          bool   `json:"is_erroneous,omitempty"`
	IsWaitingForPin      bool   `json:"is_waiting_for_pin,omitempty"`
	IsWaitingForResponse bool   `json:"is_waiting_for_response,omitempty"`
	Message              string `json:"message,omitempty"`
}
