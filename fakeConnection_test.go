package fiGo_test

import (
	"encoding/json"
	"testing"

	"github.com/Jeffail/gabs"
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
	failOnError(t, err)

	// try to get the recovery-password
	var responseAsMap map[string]string
	err = json.Unmarshal(responseAsBytes, &responseAsMap)
	failOnError(t, err)
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
	failOnError(t, err)

	// try to get informations
	var userAsMap map[string]interface{}
	err = json.Unmarshal(userByte, &userAsMap)
	failOnError(t, err)

	// check informations
	if userAsMap["access_token"] == "" || userAsMap["expires_in"] == 0 ||
		userAsMap["refresh_token"] == "" || userAsMap["scope"] == nil || userAsMap["token_type"] == nil {
		t.Fail()
	}
}

// TestAddAccount
// 1. Create a new user
// 2. Login new TestUser
// 3. Add account
// -> Check task_token
func TestAddAccount(t *testing.T) {
	username := "test@test.de"
	password := "mySuperSecretPassword"

	// create a new fakeConnection
	fakeConnection := fiGo.NewFakeConnection()
	// "store" a user
	fakeConnection.CreateUser("TestUser", username, password)
	// login
	userByte, _ := fakeConnection.CredentialLogin(username, password)
	jsonParsed, _ := gabs.ParseJSON(userByte)
	accessToken, _ := jsonParsed.Path("access_token").Data().(string)

	// add account
	responseByte, err := fakeConnection.SetupNewBankAccount(accessToken, "90090042", "de", []string{"demo", "demo"})
	failOnError(t, err)

	jsonParsed, err = gabs.ParseJSON(responseByte)
	taskToken, ok := jsonParsed.Path("task_token").Data().(string)
	if taskToken == "" || !ok || err != nil {
		t.Fail()
	}
}

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
