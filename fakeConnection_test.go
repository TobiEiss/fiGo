package fiGo_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/TobiEiss/fiGo"
)

// TestFakeConnectionImplementsEverything test only if the code compiles.
// If the code doesnt compile, the fakeConnection miss a function to implement
func TestFakeConnectionImplementsEverything(t *testing.T) {
	var connection fiGo.IConnection
	fakeConnection := fiGo.NewFakeConnection()
	connection = fakeConnection
	if connection == nil {
		t.Fail()
	}
}

// TestStoreUser creates a new user in fakeConnection
func TestStoreUser(t *testing.T) {
	// create a new fakeConnection
	fakeConnection := fiGo.NewFakeConnection()

	// "store" a user
	responseAsBytes, err := fakeConnection.CreateUser("TestUser", "test@test.de", "mySuperSecretPassword")
	if err != nil {
		t.Fail()
	}

	// try to get the recovery-password
	var responseAsMap map[string]string
	err = json.Unmarshal(responseAsBytes, &responseAsMap)
	if err != nil {
		t.Fail()
	}
}

// TestCredentialLogin tests a login.
// 1. Create a new user
// 2. Login new TestUser
// 3. Check the informations
func TestCredentialLogin(t *testing.T) {
	username := "test@test.de"
	password := "mySuperSecretPassword"

	// create a new fakeConnection
	fakeConnection := fiGo.NewFakeConnection()
	// "store" a user
	fakeConnection.CreateUser("TestUser", username, password)

	// login
	userByte, err := fakeConnection.CredentialLogin(username, password)
	if err != nil {
		t.Fail()
	}

	// try to get informations
	var userAsMap map[string]interface{}
	err = json.Unmarshal(userByte, &userAsMap)
	log.Println(string(userByte))
	if err != nil {
		t.Fail()
	}

	// check informations
	if userAsMap["access_token"] == "" || userAsMap["expires_in"] == 0 ||
		userAsMap["refresh_token"] == "" || userAsMap["scope"] == nil || userAsMap["token_type"] == nil {
		t.Fail()
	}
}
