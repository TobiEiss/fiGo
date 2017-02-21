package fiGo_test

import "testing"
import "github.com/TobiEiss/fiGo"
import "encoding/json"

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
