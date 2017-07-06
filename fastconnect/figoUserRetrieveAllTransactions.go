package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// RetrieveAllTransactions retrieves all transacions of all accounts
func RetrieveAllTransactions(connection fiGo.IConnection, accessToken string, options ...fiGo.TransactionOption) ([]interface{}, error) {
	var transactions []interface{}

	// get transactions
	answerByte, err := connection.RetrieveTransactionsOfAllAccounts(accessToken, options...)
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
