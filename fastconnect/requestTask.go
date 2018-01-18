package fastconnect

import (
	"encoding/json"

	"github.com/TobiEiss/fiGo"
)

// RequestTask check the status of a task token
func RequestTask(connection fiGo.IConnection, accessToken, taskToken string) (Task, error) {
	RequestTaskWithPinChallenge(connection, accessToken, taskToken, "", false)
}

// RequestTaskWithPinChallenge can respond to a pin challenge
func RequestTaskWithPinChallenge(connection fiGo.IConnection, accessToken, taskToken, pin string, savePin bool) (Task, error) {
	var task Task

	// try to get state of task
	answerByte, err := connection.RequestForTask(accessToken, taskToken, pin, savePin)
	if err != nil {
		return task, err
	}

	// unmarshal to Task
	err = json.Unmarshal(answerByte, &task)
	if err != nil {
		return task, err
	}

	return task, nil
}
