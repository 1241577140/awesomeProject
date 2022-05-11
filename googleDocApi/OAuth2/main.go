package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"net/url"
	"os"

	"golang.org/x/oauth2"
)

func main() {
	fmt.Println(oauthClient())
}

// oauthClient shows how to use an OAuth client ID to authenticate as an end-user.
func oauthClient() error {
	ctx := context.Background()

	// Please make sure the redirect URL is the same as the one you specified when you
	// created the client ID.
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:3080"
	}
	config := &oauth2.Config{
		ClientID:     "487707921124-isu7q87rfeanp3d204ql95lbo4skdjl4.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	// Dummy authorization flow to read auth code from stdin.
	//authURL := config.AuthCodeURL("your state")
	//fmt.Printf("Follow the link in your browser to obtain auth code:\n %s \n", authURL)

	// Read the authentication code from the command line
	u, _ := url.Parse("http://localhost:3080/?state=your+state&code=4%2F0AX4XfWgOQN4OVSt-0Iko78Iil3XI1J4OGA4HqMXHnkSEm61f4u19LlGOcoM710XEXVcjMA&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+openid&authuser=0&prompt=consent")
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m["code"][0])
	var code string
	code = m["code"][0]
	//fmt.Scanln(&code)
	code = "4%2F0AX4XfWgOQN4OVSt-0Iko78Iil3XI1J4OGA4HqMXHnkSEm61f4u19LlGOcoM710XEXVcjMA"
	// Exchange auth code for OAuth token.
	token, err := config.Exchange(ctx, code)
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
