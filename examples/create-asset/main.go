// Example: create a host asset and bind it to a node.
//
// Run with:
//
//	JUMPSERVER_URL=... JUMPSERVER_KEY_ID=... JUMPSERVER_SECRET_ID=... \
//	    JUMPSERVER_NODE_ID=... go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

func main() {
	url := os.Getenv("JUMPSERVER_URL")
	key := os.Getenv("JUMPSERVER_KEY_ID")
	secret := os.Getenv("JUMPSERVER_SECRET_ID")
	nodeID := os.Getenv("JUMPSERVER_NODE_ID")

	client := jumpserver.NewClient(
		jumpserver.WithBaseURL(url),
		jumpserver.WithAccessKeyAuth(key, secret),
	)

	ctx := context.Background()

	asset, _, err := client.Hosts.Create(ctx, &model.AssetRequest{
		Name:      "web-01",
		Address:   "192.168.1.10",
		Platform:  1,
		Protocols: []model.NamePort{{Name: "ssh", Port: 22}},
		Nodes:     []string{nodeID},
	})
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	fmt.Printf("Created asset %s (%s)\n", asset.ID, asset.Name)
}
