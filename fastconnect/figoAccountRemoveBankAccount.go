package fastconnect

import "github.com/TobiEiss/fiGo"

// RemoveBankAccount removes a bankAccount from a user on figo-side
func RemoveBankAccount(connection fiGo.IConnection, accessToken string, bankAccountID string) error {
	// deletes an user
	answerByte, err := connection.RemoveBankAccount(accessToken, bankAccountID)
	if err != nil && answerByte != nil {
		return err
	}
	return nil
}
