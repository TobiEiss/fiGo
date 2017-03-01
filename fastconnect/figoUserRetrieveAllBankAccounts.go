package fastconnect

import (
	"log"

	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// RetrieveAllBankAccounts retrieves all bankAccounts
func RetrieveAllBankAccounts(connection fiGo.IConnection, accessToken string) (interface{}, error) {
	// get accounts
	answerByte, err := connection.RetrieveAllBankAccounts(accessToken)
	if err != nil {
		return nil, err
	}

	// try to get accessToken
	log.Println(string(answerByte))
	jsonParsed, err := gabs.ParseJSON(answerByte)
	accounts, ok := jsonParsed.Path("accounts").Data().(interface{})
	if !ok {
		return accounts, err
	}
	return accounts, nil
}
