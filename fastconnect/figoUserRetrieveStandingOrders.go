package fastconnect

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/fiGo"
)

// RetrieveStandingOrders retrieves all standing orders of all accounts
func RetrieveStandingOrders(connection fiGo.IConnection, accessToken string, options ...fiGo.TransactionOption) ([]interface{}, error) {
	var standingOrders []interface{}

	// get transactions
	answerByte, err := connection.ReadStandingOrder(accessToken, options...)
	if err != nil {
		return standingOrders, err
	}

	// try to get accessToken
	jsonParsed, err := gabs.ParseJSON(answerByte)
	standingOrders, ok := jsonParsed.Search("standing_orders").Data().([]interface{})
	if !ok {
		return standingOrders, err
	}
	return standingOrders, nil
}
