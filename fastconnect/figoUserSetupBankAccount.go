package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// SetupNewBankAccount add a bankAccount to an existing figoAccount
// -> returns a taskToken
func SetupNewBankAccount(connection fiGo.IConnection, accessToken string, bankCredentials BankCredentials, savePin bool) (string, error) {
	var taskToken string

	// add BankAccount to figoAccount
	answerByte, err := connection.SetupNewBankAccount(accessToken, bankCredentials.BankCode, bankCredentials.Country, bankCredentials.Credentials, savePin)
	if err != nil {
		return taskToken, err
	}

	// try to get recoveryPassword
	jsonParsed, err := gabs.ParseJSON(answerByte)
	taskToken, ok := jsonParsed.Path("task_token").Data().(string)
	if !ok {
		return taskToken, ErrRecoveryPassword
	}
	return taskToken, nil
}
