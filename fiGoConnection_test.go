package fiGo_test

import (
	"testing"

	"github.com/TobiEiss/fiGo"
)

func TestConnection(t *testing.T) {
	connection := fiGo.NewFigoConnection("CaESKmC8MAhNpDe5rvmWnSkRE_7pkkVIIgMwclgzGcQY", "STdzfv0GXtEj_bwYn7AgCVszN1kKq5BdgEIKOM_fzybQ")
	_, err := connection.CreateUser("testUsername", "test@test.de", "mysecretpassword")

	if err != fiGo.ErrHTTPUnauthorized {
		t.Fail()
	}
}
