package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// oauthClient shows how to use an OAuth client ID to authenticate as an end-user.
func oauthClient() error {
	ctx := context.Background()

	// Please make sure the redirect URL is the same as the one you specified when you
	// created the client ID.
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:8080"
	}
	config := &oauth2.Config{
		ClientID:     "714842259881-97khn33kah9q33e9qeql9e6m4bkqockn.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	// Dummy authorization flow to read auth code from stdin.
	authURL := config.AuthCodeURL("your state")
	fmt.Printf("Follow the link in your browser to obtain auth code: %s", authURL)

	// Read the authentication code from the command line
	//var code string
	//fmt.Scanln(&code)

	// Exchange auth code for OAuth token.
	token, err := config.Exchange(ctx, "4%2F0AX4XfWiqI2Tz1IpItvUlCmaLo6qMAakAH4TSwkalLo7-YkZfbh5-5BFyXqIx9caAl9_EiA")
	if err != nil {
		return fmt.Errorf("config.Exchange: %v", err)
	}
	client, err := pubsub.NewClient(ctx, "your-project-id", option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	// Use the authenticated client.
	_ = client

	return nil
}
func main() {
	oauthClient()
}
