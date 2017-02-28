package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// LoginUser logs a user in
// -> returns the accessToken
func LoginUser(connection fiGo.IConnection, figoUser FigoUser) (string, error) {
	var accessToken string

	// try to login
	answerByte, err := connection.CredentialLogin(figoUser.Email, figoUser.Password)
	if err != nil {
		return accessToken, err
	}

	// try to get accessToken
	jsonParsed, err := gabs.ParseJSON(answerByte)
	accessToken, ok := jsonParsed.Path("access_token").Data().(string)
	if !ok {
		return accessToken, ErrRecoveryPassword
	}
	return accessToken, nil
}
