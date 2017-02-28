package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// CreateUser creates a user on figo-side
// -> return a recovery-password
func CreateUser(connection fiGo.IConnection, figoUser FigoUser) (string, error) {
	var recoveryPassword string

	// setup new figo-user
	answerByte, err := connection.CreateUser(figoUser.Username, figoUser.Email, figoUser.Password)
	if err != nil {
		return recoveryPassword, err
	}

	// try to get recoveryPassword
	jsonParsed, err := gabs.ParseJSON(answerByte)
	recoveryPassword, ok := jsonParsed.Path("recovery_password").Data().(string)
	if !ok {
		return recoveryPassword, ErrRecoveryPassword
	}
	return recoveryPassword, nil
}
