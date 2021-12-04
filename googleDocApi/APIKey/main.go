package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// apiKey shows how to use an API key to authenticate.
func apiKey() error {
	client, err := pubsub.NewClient(context.Background(), "inner-geography-333815", option.WithAPIKey("AIzaSyCj_dHvBGDYgzVXiX0eIqwLDIrl6npTwYw"))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	// Use the authenticated client.
	_ = client

	return nil
}
func main() {
	apiKey()
}
