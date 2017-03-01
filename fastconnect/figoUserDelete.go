package fastconnect

import "github.com/TobiEiss/fiGo"

// DeleteUser deletes a user on figo-side
func DeleteUser(connection fiGo.IConnection, accessToken string) error {
	// deletes an user
	answerByte, err := connection.DeleteUser(accessToken)
	if err != nil && answerByte != nil {
		return err
	}
	return nil
}
