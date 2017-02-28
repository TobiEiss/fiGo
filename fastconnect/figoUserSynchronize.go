package fastconnect

import (
	"github.com/TobiEiss/fiGo"
)

// SynchronizeFigoUser starts a synchronize-Task on figo-Side
func SynchronizeFigoUser(connection fiGo.IConnection, accessToken string, taskToken string) error {
	_, err := connection.CreateNewSynchronizationTask(accessToken, taskToken)
	return err
}
