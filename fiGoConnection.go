package fiGo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jeffail/gabs"
)

const (
	baseURL             = "https://api.figo.me"
	authUserURL         = "/auth/user"
	authTokenURL        = "/auth/token"
	restAccountsURL     = "/rest/accounts"
	restUserURL         = "/rest/user"
	restTransactionsURL = "/rest/transactions"
)

var (
	// ErrUserAlreadyExists - code 30002
	ErrUserAlreadyExists = errors.New("user already exists")

	// ErrHTTPUnauthorized - code 90000
	ErrHTTPUnauthorized = errors.New("invalid authorization")
)

// IConnection represent an interface for connections.
// This provides to use fakeConnection and a real-figoConnection
type IConnection interface {
	// http://docs.figo.io/#create-new-figo-user
	// Ask user for new name, email and password
	CreateUser(name string, email string, password string) ([]byte, error)

	// http://docs.figo.io/#credential-login
	// Login with email (aka username) and password
	CredentialLogin(username string, password string) ([]byte, error)

	// http://docs.figo.io/#delete-current-user
	// Remove user with an accessToken. You get the token after successfully login
	DeleteUser(accessToken string) ([]byte, error)

	// http://docs.figo.io/#setup-new-bank-account
	// Add a BankAccount to figo-Account
	// -> yout get accessToken from the login-response
	// -> country is something like "de" (for germany)
	SetupNewBankAccount(accessToken string, bankCode string, country string, credentials []string) ([]byte, error)

	// http://docs.figo.io/#retrieve-transactions-of-one-or-all-account
	// Retrieves all Transactions
	RetrieveTransactionsOfAllAccounts(accessToken string) ([]byte, error)
}

// Connection represent a connection to figo
type Connection struct {
	AuthString string
}

// NewFigoConnection creates a new connection.
// -> You need a clientID and a clientSecret. (You will get this from figo.io)
func NewFigoConnection(clientID string, clientSecret string) *Connection {
	authInfo := clientID + ":" + clientSecret
	authString := "Basic " + base64.URLEncoding.EncodeToString([]byte(authInfo))
	return &Connection{AuthString: authString}
}

// CreateUser creates a new user.
// Ask User for name, email and a password
func (connection *Connection) CreateUser(name string, email string, password string) ([]byte, error) {
	// build url
	url := baseURL + authUserURL

	// build jsonBody
	requestBody := map[string]string{
		"name":     name,
		"email":    email,
		"password": password}
	jsonBody, err := json.Marshal(requestBody)

	// build request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, connection.AuthString)
}

// CredentialLogin create a login.
// -> First you have to create a user (CreateUser)
func (connection *Connection) CredentialLogin(username string, password string) ([]byte, error) {
	// build url
	url := baseURL + authTokenURL

	// build jsonBody
	requestBody := map[string]string{
		"grant_type": "password",
		"username":   username,
		"password":   password}
	jsonBody, err := json.Marshal(requestBody)

	// build request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, connection.AuthString)
}

// SetupNewBankAccount add a new bankAccount to an existing figo-Account
func (connection *Connection) SetupNewBankAccount(accessToken string, bankCode string, country string, credentials []string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := baseURL + restAccountsURL

	// build jsonBody
	requestBody := map[string]interface{}{
		"bank_code":   bankCode,
		"country":     country,
		"credentials": credentials,
	}
	jsonBody, err := json.Marshal(requestBody)

	// build request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// DeleteUser deletes an existing user
func (connection *Connection) DeleteUser(accessToken string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := baseURL + restUserURL

	// build request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// RetrieveTransactionsOfAllAccounts with accessToken from login-session
func (connection *Connection) RetrieveTransactionsOfAllAccounts(accessToken string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := baseURL + restTransactionsURL

	// build request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

func buildRequestAndCheckResponse(request *http.Request, authString string) ([]byte, error) {
	// set headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authString)

	// setup client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// get response
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	// check response for errors
	if string(body) != "" {
		jsonParsed, err := gabs.ParseJSON(body)
		if err != nil {
			return nil, err
		}
		value, ok := jsonParsed.Path("error.code").Data().(float64)
		if ok {
			if value == 30002 {
				return body, ErrUserAlreadyExists
			} else if value == 90000 {
				return body, ErrHTTPUnauthorized
			}
		}
	}

	return body, nil
}
