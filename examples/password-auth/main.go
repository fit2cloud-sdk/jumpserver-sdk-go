package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
)

func main() {
	base_url := os.Getenv("JUMPSERVER_URL")
	username := os.Getenv("JUMPSERVER_USERNAME")
	password := os.Getenv("JUMPSERVER_PASSWORD")
	fmt.Println(base_url, username)
	// Create a client using password authentication.
	// This will automatically obtain a Bearer token from the
	// JumpServer authentication endpoint.
	client := jumpserver.NewClient(
		jumpserver.WithBaseURL(base_url),
		jumpserver.WithPasswordAuth(username, password),
	)

	// The client will automatically authenticate with a Bearer token.
	// No need to manually obtain the token first.
	user, resp, err := client.Users.Profile(context.Background())
	if err != nil {
		log.Fatalf("Failed to get user profile: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("User: %+v\n", user)
}
