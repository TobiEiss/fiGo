package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// RetrieveAllTransactions retrieves all transacions of all accounts
func RetrieveAllTransactions(connection fiGo.IConnection, accessToken string) ([]map[string]interface{}, error) {
	var transactions []map[string]interface{}

	// get transactions
	answerByte, err := connection.RetrieveTransactionsOfAllAccounts(accessToken)
	if err != nil {
		return transactions, err
	}

	// try to get accessToken
	jsonParsed, err := gabs.ParseJSON(answerByte)
	transactions, ok := jsonParsed.Path("transactions").Data().([]map[string]interface{})
	if !ok {
		return transactions, err
	}
	return transactions, nil
}
