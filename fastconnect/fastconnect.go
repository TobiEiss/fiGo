package fastconnect

import "errors"

var (
	// ErrRecoveryPassword cant find a "recovery_password" in JSON
	ErrRecoveryPassword = errors.New("figo-fastconnect: extract recovery password failed")
)
