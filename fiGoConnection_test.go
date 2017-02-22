package fiGo_test

import (
	"testing"

	"github.com/TobiEiss/fiGo"
)

func TestConnection(t *testing.T) {
	var connection fiGo.IConnection
	figoConnection := fiGo.NewFigoConnection("CaESKmC8MAhNpDe5rvmWnSkRE_7pkkVIIgMwclgzGcQY", "STdzfv0GXtEj_bwYn7AgCVszN1kKq5BdgEIKOM_fzybQ")
	connection = figoConnection
	_, err := connection.CreateUser("testUsername", "test@test.de", "mysecretpassword")

	if err != fiGo.ErrHTTPUnauthorized {
		t.Fail()
	}
}

func TestLogin(t *testing.T) {
	var connection fiGo.IConnection
	figoConnection := fiGo.NewFigoConnection("CaESKmC8MAhNpDe5rvmWnSkRE_7pkkVIIgMwclgzGcQY", "STdzfv0GXtEj_bwYn7AgCVszN1kKq5BdgEIKOM_fzybQ")
	connection = figoConnection
	_, err := connection.CredentialLogin("demo@figo.me", "demo1234")

	if err != fiGo.ErrHTTPUnauthorized {
		t.Fail()
	}
}
