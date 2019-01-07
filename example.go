package main

import (
	"./api"
	"context"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	conf *oauth2.Config
	ctx  context.Context
)

func printProfiles(client *http.Client) {
	bufferClient := api.NewClient(client)
	profiles, err := bufferClient.Profile.ListProfiles()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Profiles:", profiles)
}

func main() {
	ctx = context.Background()
	var clientSecret string
	flag.StringVar(&clientSecret, "clientSecret", "", "Your client secret")
	flag.Parse()

	if clientSecret == "" {
		log.Fatal(errors.New("Missing flag: clientSecret. Start this example with -clientSecret"))
	}

	conf = &oauth2.Config{
		ClientID:     "5c32a3174b1be71532161683", // Change to your client id: https://buffer.com/developers/apps/
		ClientSecret: clientSecret,
		Scopes:       nil,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://bufferapp.com/oauth2/authorize",
			TokenURL: "https://api.bufferapp.com/1/oauth2/token.json",
		},
		// my own callback URL
		RedirectURL: "http://127.0.0.1:9999/oauth/callback",
	}

	ctx = context.Background()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	authUrl := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)

	log.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)
	err := open.Run(authUrl)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", authUrl)

	http.HandleFunc("/oauth/callback", oAuthCallbackHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))

}

/**
Method to handle OAuth callback, not library specific
*/
func oAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	// Use the authorization code that is pushed to the redirect
	// URL.
	code := queryParts["code"][0]
	log.Printf("code: %s\n", code)

	// Exchange will do the handshake to retrieve the initial access token.
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token: %s", tok)
	// The HTTP Client returned by conf.Client will refresh the token as necessary.
	client := conf.Client(ctx, tok)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(color.CyanString("Authentication successful"))
	}

	// show succes page
	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)

	printProfiles(client)
}
