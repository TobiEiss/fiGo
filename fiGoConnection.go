package fiGo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs"
)

const (
	defaultBaseURL        = "https://api.figo.me"
	authUserURL           = "/auth/user"
	authTokenURL          = "/auth/token"
	restAccountsURL       = "/rest/accounts"
	restUserURL           = "/rest/user"
	restTransactionsURL   = "/rest/transactions"
	restStandingOrdersURL = "/rest/standing_orders"
	restSyncURL           = "/rest/sync"
	taskProgressURL       = "/task/progress"
	catalog               = "/catalog"
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
	// Set another host
	SetHost(host string)

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
	// -> you get accessToken from the login-response
	// -> country is something like "de" (for germany)
	SetupNewBankAccount(accessToken string, bankCode string, country string, credentials []string) ([]byte, error)

	// http://docs.figo.io/#retrieve-all-bank-accounts
	// Retrieves all bankAccounts for an user
	RetrieveAllBankAccounts(accessToken string) ([]byte, error)

	// http://docs.figo.io/#delete-bank-account
	// Removes a bankAccount from figo-account
	RemoveBankAccount(accessToken string, bankAccountID string) ([]byte, error)

	// http://docs.figo.io/#poll-task-state
	// request a task
	// -> you need a taskToken. You will get this from SetupNewBankAccount
	RequestForTask(accessToken string, taskToken string) ([]byte, error)

	// http://docs.figo.io/#retrieve-transactions-of-one-or-all-account
	// Retrieves all Transactions
	RetrieveTransactionsOfAllAccounts(accessToken string, options ...TransactionOption) ([]byte, error)

	// http://docs.figo.io/#retrieve-a-transaction
	// Retrieves a specific Transaction
	RetrieveSpecificTransaction(accessToken string, transactionID string) ([]byte, error)

	// http://docs.figo.io/#tag/Catalog
	// Read individual Catalog Entry
	ReadIndividualCatalogEntry(accessToken string, catalogCategory string, countryCode string, serviceID string) ([]byte, error)

	// http://docs.figo.io/#standing-orders-api-calls-read-standing-orders
	// Read standing orders
	ReadStandingOrder(accessToken string, options ...TransactionOption) ([]byte, error)
}

// Connection represent a connection to figo
type Connection struct {
	AuthString string
	Host       string
}

// NewFigoConnection creates a new connection.
// -> You need a clientID and a clientSecret. (You will get this from figo.io)
func NewFigoConnection(clientID string, clientSecret string) *Connection {
	authInfo := clientID + ":" + clientSecret
	authString := "Basic " + base64.URLEncoding.EncodeToString([]byte(authInfo))
	return &Connection{AuthString: authString, Host: defaultBaseURL}
}

// SetHost sets a new host
func (connection *Connection) SetHost(host string) {
	connection.Host = host
}

// CreateUser creates a new user.
// Ask User for name, email and a password
func (connection *Connection) CreateUser(name string, email string, password string) ([]byte, error) {
	// build url
	url := connection.Host + authUserURL

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
	url := connection.Host + authTokenURL

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
	url := connection.Host + restAccountsURL

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

// RemoveBankAccount removes a bank account
func (connection *Connection) RemoveBankAccount(accessToken string, bankAccountID string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + restAccountsURL + "/" + bankAccountID

	// build request
	request, err := http.NewRequest("DELETE", url, nil)
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
	url := connection.Host + restUserURL

	// build request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// RequestForTask starts a new task to synchronize real bankAccount and figoAccount
func (connection *Connection) RequestForTask(accessToken string, taskToken string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + taskProgressURL + "?id=" + taskToken

	// build request
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// SyncTask represents a synchronization task
type SyncTask struct {
	State                string   `json:"state"`
	RedirectURI          string   `json:"redirect_uri"`
	DisableNotifications bool     `json:"disable_notifications"`
	IfNotSyncedSince     int      `json:"if_not_synced_since"`
	AutoContinue         bool     `json:"auto_continue"`
	AccountIds           []string `json:"account_ids"`
	SyncTasks            []string `json:"sync_tasks"`
}

// CreateSynchronizationTask creates a new task to synchronize
func (connection *Connection) CreateSynchronizationTask(accessToken string, syncTask SyncTask) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + restSyncURL

	marshaledSyncTask, err := json.Marshal(&syncTask)
	if err != nil {
		return nil, err
	}

	// build request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(marshaledSyncTask))
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// TransactionOption are options for transaction-calls
type TransactionOption struct {
	Key   TransactionOptionKey
	Value string
}

// TransactionOptionKey are the type for allowed keys
type TransactionOptionKey string

var (
	// Cent if true, the amount of the transactions will be shown in cents.
	Cent TransactionOptionKey = "cents"
)

// RetrieveTransactionsOfAllAccounts with accessToken from login-session
func (connection *Connection) RetrieveTransactionsOfAllAccounts(accessToken string, options ...TransactionOption) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + restTransactionsURL

	// build request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// if there are options, apply these
	if len(options) > 0 {
		q := request.URL.Query()
		for _, option := range options {
			q.Add(string(option.Key), option.Value)
		}
		request.URL.RawQuery = q.Encode()
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// ReadStandingOrder all standing orders
func (connection *Connection) ReadStandingOrder(accessToken string, options ...TransactionOption) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + restStandingOrdersURL

	// build request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// if there are options, apply these
	if len(options) > 0 {
		q := request.URL.Query()
		for _, option := range options {
			q.Add(string(option.Key), option.Value)
		}
		request.URL.RawQuery = q.Encode()
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// RetrieveSpecificTransaction with accessToken from login-session
func (connection *Connection) RetrieveSpecificTransaction(accessToken string, transactionID string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + restTransactionsURL + "/" + transactionID

	// build request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// RetrieveAllBankAccounts retrieves all bankAccounts for an user
func (connection *Connection) RetrieveAllBankAccounts(accessToken string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + restAccountsURL

	// build request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return buildRequestAndCheckResponse(request, accessToken)
}

// ReadIndividualCatalogEntry gets all infos for a specific catalog entry.
// catalogCategory is "banks" or "services"
// countryCode is "de" or "at"
// serviceID is something like BLZ
func (connection *Connection) ReadIndividualCatalogEntry(accessToken string, catalogCategory string, countryCode string, serviceID string) ([]byte, error) {
	// build accessToken
	accessToken = "Bearer " + accessToken

	// build url
	url := connection.Host + catalog + "/" + catalogCategory + "/" + serviceID

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
