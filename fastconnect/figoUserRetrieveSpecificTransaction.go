package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// RetrieveSpecificTransactions retrieves a specific transaction
func RetrieveSpecificTransactions(connection fiGo.IConnection, accessToken string) ([]interface{}, error) {
	var transactions []interface{}

	// get transactions
	answerByte, err := connection.RetrieveTransactionsOfAllAccounts(accessToken)
	if err != nil {
		return transactions, err
	}

	// try to get accessToken
	jsonParsed, err := gabs.ParseJSON(answerByte)
	transactions, ok := jsonParsed.Search("transactions").Data().([]interface{})
	if !ok {
		return transactions, err
	}
	return transactions, nil
}
