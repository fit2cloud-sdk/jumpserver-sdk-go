// Integration test: full CRUD lifecycle for every JumpServer v4 service.
//
// Environment variables required:
//
//	JUMPSERVER_URL        — base URL
//	JUMPSERVER_KEY_ID     — access key ID
//	JUMPSERVER_SECRET_ID  — access key secret
//
// Run:
//
//	go run ./examples/integration
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
)

var (
	passed  int
	failed  int
	skipped int

	client *jumpserver.Client
	scoped *jumpserver.Client
	ctx    context.Context
	ts     string
	orgID  string

	createdNodeID      string
	createdZoneID      string
	createdLabelID     string
	createdHostID      string
	createdAccountID   string
	createdTemplateID  string
	createdPermID      string
	createdCmdFilterID string
	createdCmdGroupID  string
	createdCategoryIDs = make(map[string]string)
)

func ok(name string) {
	fmt.Printf("  \xe2\x9c\x93 %s\n", name)
	passed++
}

func fail(name string, err error) {
	fmt.Printf("  \xe2\x9c\x97 %-50s %s\n", name, shortErr(err))
	failed++
}

func skip(name, reason string) {
	fmt.Printf("  \xe2\x97\x8b %-50s [SKIP: %s]\n", name, reason)
	skipped++
}

func shortErr(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	if len(s) > 160 {
		return s[:160] + "..."
	}
	return s
}

func section(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	url := os.Getenv("JUMPSERVER_URL")
	keyID := os.Getenv("JUMPSERVER_KEY_ID")
	secretID := os.Getenv("JUMPSERVER_SECRET_ID")
	if url == "" || keyID == "" || secretID == "" {
		fmt.Fprintln(os.Stderr, "JUMPSERVER_URL, JUMPSERVER_KEY_ID, JUMPSERVER_SECRET_ID required")
		os.Exit(1)
	}

	client = jumpserver.NewClient(
		jumpserver.WithBaseURL(url),
		jumpserver.WithAccessKeyAuth(keyID, secretID),
	)
	ctx = context.Background()
	ts = fmt.Sprintf("%d", time.Now().UnixNano()%100000)

	orgs, _, err := client.Orgs.List(ctx, &jumpserver.ListOptions{Limit: 10})
	if err == nil && len(orgs) > 0 {
		for _, o := range orgs {
			if !o.IsRoot {
				orgID = o.ID
				break
			}
		}
		if orgID == "" {
			orgID = orgs[0].ID
		}
	}
	if orgID == "" {
		fmt.Fprintln(os.Stderr, "No organizations found")
		os.Exit(1)
	}
	fmt.Printf("Using org: %s\n", orgID)
	scoped = client.WithOrgScope(orgID)

	testSettings()
	testUsers()
	testRoles()
	testOrgs()
	testPlatforms()
	testNodes()
	testZones()
	testLabels()
	testAssets()
	testAccountTemplates()
	testAccounts()
	testChangeSecrets()
	testAccountBackups()
	testPermissions()
	testSelfAssets()
	testPermRelations()
	testCommandFilters()
	testCommandGroups()
	testLoginACLs()
	testTickets()
	testAudits()
	testTerminal()
	testXpack()
	testOps()
	testWalkPages()

	cleanup()

	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Results: %d passed, %d failed, %d skipped, %d total\n",
		passed, failed, skipped, passed+failed+skipped)
	if failed > 0 {
		os.Exit(1)
	}
}
