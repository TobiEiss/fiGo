package fastconnect

import (
	"encoding/json"

	"github.com/TobiEiss/fiGo"
)

// RequestTask ask for task
func RequestTask(connection fiGo.IConnection, accessToken string, taskToken string) (Task, error) {
	var task Task

	// try to get state of task
	answerByte, err := connection.RequestForTask(accessToken, taskToken)
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
