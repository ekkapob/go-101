package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type Data struct {
	Accounts []Account
}

type Account struct {
	Username string
	Password string
	Scopes   []string
	Role     string
}

func readAccounts(filename string) Data {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("unable to open data file")
	}
	var accountData Data
	err = json.Unmarshal(data, &accountData)
	if err != nil {
		log.Fatal("unable to unmarshal data ", err)
	}
	return accountData
}

func (d Data) GetAccount(username, password string) (*Account, error) {
	errMsg := "invalid username or password"
	if username == "" || password == "" {
		return nil, errors.New(errMsg)
	}
	for _, account := range d.Accounts {
		if account.Username == username && account.Password == password {
			return &account, nil
		}
	}
	return nil, errors.New(errMsg)
}
