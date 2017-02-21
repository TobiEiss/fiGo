package fiGo

import (
	"encoding/json"
)

// FakeConnection is a fakes the figoAPI
type FakeConnection struct {
	Users []map[string]string
}

// NewFakeConnection creates a new "fak-Connection" only in memory
func NewFakeConnection() *FakeConnection {
	return &FakeConnection{Users: make([]map[string]string, 0)}
}

// CreateUser "store" a user in this fake-Connection
func (fakeConnection *FakeConnection) CreateUser(name string, email string, password string) ([]byte, error) {
	recoveryPassword := "recoveryPassword"
	// "store" in fakeConnection
	fakeConnection.Users = append(fakeConnection.Users, map[string]string{
		"name":              name,
		"email":             email,
		"password":          password,
		"recovery_password": recoveryPassword})

	responseMap := map[string]string{"recovery_password": recoveryPassword}
	return json.Marshal(responseMap)
}
