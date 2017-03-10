# fiGo - a go driver for the figo-API (figo.io)

[![Build Status](https://travis-ci.org/TobiEiss/fiGo.svg?branch=master)](https://travis-ci.org/TobiEiss/fiGo)

This project is a golang-driver for [figo](http://www.figo.io).  
If you want to use this, you need a clientID and a clientSecret. You will get this from figo.

You miss something here? - Please let me know!

Currently implemented:
* [create a user](#create-a-user) ([figo-API-reference](http://docs.figo.io/#create-new-figo-user))
* [credential login](#credentials-login) ([figo-API-reference](http://docs.figo.io/#credential-login))
* [setup new bank account](#setup-new-bank-account) ([figo-API-reference](http://docs.figo.io/#setup-new-bank-account))
* [delete a user](#delete-a-user) ([figo-API-reference](http://docs.figo.io/#delete-a-user))
* [retrieve transactions and account-informations](#retrieve-transactions-and-account-informations)

## Getting started

Install fiGo:
`go get github.com/TobiEiss/fiGo`

Dependencies:
- [gabs](https://github.com/Jeffail/gabs) for parsing, creating and editing unknown or dynamic JSON in golang

## Usage

First create a new connection:
```golang
// create a new fiGo-Connection
var connection fiGo.IConnection
figoConnection := fiGo.NewFigoConnection("clientID", "clientSecret")
connection = figoConnection
```

### Create a user

Ask your user for an username, email-address and password. Then add to figo:
```golang
recPwByteArray, err := connection.CreateUser("testUsername", "test@test.de", "mysecretpassword")
if err == fiGo.ErrHTTPUnauthorized {
    // TODO: handle if this was unauthorized
} else if err == fiGo.ErrUserAlreadyExists {
    // TODO: handle if user already exists
}
```

You will get back a recovery-password in JSON-format as byte array:
```json
{"recovery_password": "abcd-efgh-ijkl-mnop"}
```

Fast way to get this (use [gabs](https://github.com/Jeffail/gabs)):
```golang
jsonParsed, err := gabs.ParseJSON(recPwByteArray)
recoveryPassword, ok := jsonParsed.Path("recovery_password").Data().(string)
if ok {
    // do whatever you want with the "recoveryPassword"
}
```

### Credentials login

Login your users:

```golang
userAsJson, err := connection.CredentialLogin("test@test.de", "mysecretpassword")
// TODO error handling
```

You will get all relevant user data like this:
```json
{
   "access_token":"abcdefghijklmnopqrstuvwxyz",
   "token_type":"Bearer",
   "expires_in":600.0,
   "refresh_token":"abcdefghijklmnopqrstuvwxyz",
   "scope":"accounts=rw transactions=rw balance=rw user=rw offline create_user "
}
```

Tip: Use [gabs](https://github.com/Jeffail/gabs) to get specific fields.  
Notice: Keep the `access_token` for other user-activities.

### Setup new bank account

Add a bankAccount to an existing figo-account

```golang
jsonAnswer, err := connection.SetupNewBankAccount(value, "90090042", "de", []string{"demo", "demo"})
```

The `jsonAnswer` contains a `task_token`. You need this to sync the figo-account with a real bank-account.
```json
{"task_token": "abcdefghijklmnopqrstuvwxyz"}
```

### Delete a user

You want to delete a user? - No problem. Just call code below:
```golang
jsonAnswer, err := connection.DeleteUser(accessToken)
```

### Retrieve transactions and account-informations

To retrieve transactions use the access-Token from [credential login](#credentials-login):
```golang
answerByte, err := connection.RetrieveTransactionsOfAllAccounts(accessToken)
```

For Account-Information:

```golang
answerByte, err := connection.RetrieveAllBankAccounts(accessToken)
```

You will get back the transactions and account-informations as JSON. Use gabs and Json.Unmarshal to put this directly in a model.

## fastconnect

FiGo-fastconnect is a way to interact in a faster way with figo. All the user-/ account-/ .. models are ready. Also all the API-calls.

### Getting started with fastconnect

```golang
// First, let's create some struct-objects for a figoUser and bankCredentials.
figoUser := fastconnect.FigoUser{
    Email:    "email@example.com",
    Username: "username",
    Password: "mySecretPassword",
}
bankCredentials := fastconnect.BankCredentials{
    BankCode:    "90090042",
    Country:     "de",
    Credentials: []string{"demo", "demo"},
}

// Create a new connection to figo
figoConnection := fiGo.NewFigoConnection("clientID", "clientSecret")

// Now create the user on figo-side
recoveryPassword, err := fastconnect.CreateUser(figo.Connection, figoUser)
if recoveryPassword == "" || err != nil {
    // TODO: handle error!
}

// Login the new created user to get an accessToken
accessToken, err = fastconnect.LoginUser(figo.Connection, figoUser)
if accessToken == "" || err != nil {
    // Can't create figoUser. TODO: handle this!
}

// Add BankCredentials to the figo-account on figo-side
taskToken, err := fastconnect.SetupNewBankAccount(figo.Connection, accessToken, bankCredentials)
if err != nil || taskToken == "" {
    // Error while setup new bankAccount. TODO handle error!
}

// We need to check the snychronize-Task
task, err := fastconnect.RequestTask(figo.Connection, accessToken, taskToken)
if err != nil {
    // Should i say something? - Yeah..TODO: handle error!
}

// NOTICE! Check now the task, if everything is ready synchronized. If not, ask again.

// Now, you can retrieve all transations
transactionInterfaces, err := fastconnect.RetrieveAllTransactions(figo.Connection, accessToken)
if err != nil || transactionInterfaces == nil {
    // TODO: handle your error here!
}

// convert now to a model. TODO: implement a "Transaction" model with "json"-tags.
transactions := make(Transaction, 0)
for _, transactionInterface := range transactionInterfaces {
    transactionByte, err := json.Marshal(transactionInterface)
    if err == nil {
        transaction := Transaction{}
        json.Unmarshal(transactionByte, &transaction)
        transactions = append(transactions, transaction)
    }
}
```

Checkout the [fastconnect](https://github.com/TobiEiss/fiGo/fastconnect/) package for more stuff!