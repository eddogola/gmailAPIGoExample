package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"

	"golang.org/x/oauth2"
)

type Cred struct {
	Installed struct{
		Client_secret string `json:"client_secret"`
		Client_id string `json:"client_id"`
	} `json:"installed"`
}

func getClientCreds() Cred {
	filename := "credentials.json"
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var cred Cred
	json.Unmarshal(bs, &cred)

	return cred
}

func main() {
	creds := getClientCreds()

	fmt.Println("client id:", creds.Installed.Client_id)
	fmt.Println("client secret:", creds.Installed.Client_secret)
}