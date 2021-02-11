package main

import (
	"io/ioutil"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	// "google.golang.org/api/gmail/v1"
)

func check(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func getConfig() *oauth2.Config {
	filename := "credentials.json"
	credentials, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	config, err := google.ConfigFromJSON(credentials, "https://www.googleapis.com/auth/gmail.send") // gmail.GmailSendScope) //"https://www.googleapis.com/auth/gmail.send"
	check(err)

	return config
}

func main() {
	fmt.Println(getConfig())
}