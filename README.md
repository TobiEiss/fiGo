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
