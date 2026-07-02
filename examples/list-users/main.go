// Example: list users from a JumpServer instance using an AccessKey.
//
// Run with:
//
//	JUMPSERVER_URL=http://localhost JUMPSERVER_KEY_ID=xxx JUMPSERVER_SECRET_ID=yyy go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
)

func main() {
	url := os.Getenv("JUMPSERVER_URL")
	key := os.Getenv("JUMPSERVER_KEY_ID")
	secret := os.Getenv("JUMPSERVER_SECRET_ID")
	if url == "" || key == "" || secret == "" {
		log.Fatal("JUMPSERVER_URL, JUMPSERVER_KEY_ID, JUMPSERVER_SECRET_ID required")
	}

	client := jumpserver.NewClient(
		jumpserver.WithBaseURL(url),
		jumpserver.WithAccessKeyAuth(key, secret),
	)

	ctx := context.Background()

	var all []string
	err := jumpserver.WalkPages(ctx, &jumpserver.ListOptions{Limit: 100}, 0,
		func(ctx context.Context, opts *jumpserver.ListOptions) (*jumpserver.Response, error) {
			usrs, resp, err := client.Users.List(ctx, nil, opts)
			if err != nil {
				return resp, err
			}
			for _, u := range usrs {
				all = append(all, u.Username)
			}
			return resp, nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d users: %v\n", len(all), all)
}
