package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func check(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func createClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	token, err := getTokenFromFile(tokenFile)

	if err != nil {
		token = getTokenFromWeb(config)
		saveTokenToFile(tokenFile, token)
	}

	return config.Client(context.Background(), token)
}

func getConfig() *oauth2.Config {
	filename := "credentials.json"
	credentials, err := ioutil.ReadFile(filename)

	check(err)

	config, err := google.ConfigFromJSON(credentials, gmail.GmailReadonlyScope)
	check(err)

	return config
}

func getTokenFromFile(filename string) (*oauth2.Token, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)

	return token, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	// creates a URL for the user to follow
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Visit this URL and paste the authorization code: %v", url)

	// grabs the authorization code you paste into the terminal
	var code string
	_, err := fmt.Scan(&code)
	check(err)

	// exchange the code for an access token
	token, err := config.Exchange(context.TODO(), code)
	check(err)

	return token
}

// saves a token to a file path
func saveTokenToFile(path string, token *oauth2.Token) {
	fmt.Printf("saving credential file to: %v", path)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	check(err)

	defer file.Close()
	json.NewEncoder(file).Encode(token)
}

func main() {
	config := getConfig()

	client := createClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve gmail client: %v", err)
	}

	user := "me"
	resp, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(resp.Labels) == 0 {
		fmt.Println("No labels found")
		return
	}
	fmt.Println("Labels:")
	for _, label := range resp.Labels {
		fmt.Printf("- %s\n", label.Name)
	}
}