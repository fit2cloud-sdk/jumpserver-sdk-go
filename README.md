[![CN](https://img.shields.io/badge/CN-中文-red)](README_CN.md)
[![EN](https://img.shields.io/badge/EN-English-blue)](README.md)

# jumpserver-sdk-go

Go SDK for [JumpServer](https://www.jumpserver.org/) REST API, targeting **v4.10.x**.
[中文文档](README_CN.md) | [English](README.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/fit2cloud-sdk/jumpserver-sdk-go.svg)](https://pkg.go.dev/github.com/fit2cloud-sdk/jumpserver-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/fit2cloud-sdk/jumpserver-sdk-go)](https://goreportcard.com/report/github.com/fit2cloud-sdk/jumpserver-sdk-go)

## Features

- **Full CRUD coverage** — 26 service modules covering users, assets, accounts, permissions, audits, tickets, ops jobs, and more
- **Typed asset categories** — Hosts, Devices, Databases, Webs, Clouds, Customs each with dedicated CRUD operations
- **Multiple auth methods** — AccessKey (HMAC-SHA256), Bearer Token, Private Token, Password Auth (username/password), custom Authenticator
- **Organization scope** — `WithOrgScope(id)` switches org context without rebuilding the client
- **Auto pagination** — `WalkPages()` iterates through all pages automatically
- **Smart retry** — Exponential backoff with full jitter, retries only transient errors (timeout, connection reset, 429/5xx)
- **Zero third-party dependencies** — pure standard library
- **Go 1.25** — uses `math/rand/v2`, `maps.Clone`, `for range int`, and other modern features

## Requirements

- Go 1.25+

## Installation

```bash
go get github.com/fit2cloud-sdk/jumpserver-sdk-go
```

## Quick Start

```go
package main
import (
    "context"
    "fmt"
    "log"
    "os"
    jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
    "github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)
func main() {
    client := jumpserver.NewClient(
        jumpserver.WithBaseURL(os.Getenv("JUMPSERVER_URL")),
        jumpserver.WithAccessKeyAuth(
            os.Getenv("JUMPSERVER_KEY_ID"),
            os.Getenv("JUMPSERVER_SECRET_ID"),
        ),
    )
    ctx := context.Background()
    // List users
    users, _, err := client.Users.List(ctx, nil, &jumpserver.ListOptions{Limit: 20})
    if err != nil {
        log.Fatal(err)
    }
    for _, u := range users {
        fmt.Println(u.Username, u.Email)
    }
    // Filter by condition
    users, _, _ = client.Users.List(ctx,
        map[string]string{"username": "admin"},
        &jumpserver.ListOptions{Limit: 10},
    )
    // Create a host asset
    host, _, _ := client.Hosts.Create(ctx, &model.AssetRequest{
        Name:      "web-01",
        Address:   "192.168.1.10",
        Platform:  1, // Linux platform ID
        Protocols: []model.NamePort{{Name: "ssh", Port: 22}},
    })
    fmt.Println("Created:", host.ID)
}
```

## Authentication

```go
// AccessKey HMAC-SHA256 signature (recommended for service accounts)
jumpserver.WithAccessKeyAuth(keyID, secretID)
// Bearer Token
jumpserver.WithBearerToken(token)
// Private Token (Authorization: Token <token>)
jumpserver.WithPrivateToken(token)
// Password Auth (username/password, auto-obtains Bearer token)
jumpserver.WithPasswordAuth(username, password)
// Custom authenticator
jumpserver.WithAuthenticator(myAuth)
```

## Organization Scope

JumpServer routes most endpoints through organizations. The default header is `X-JMS-ORG: ROOT`.

```go
// Set default organization
client := jumpserver.NewClient(
    jumpserver.WithBaseURL(url),
    jumpserver.WithOrg("org-uuid"),
    // ...
)
// Derive a scoped client (shares underlying HTTP connection)
scoped := client.WithOrgScope("other-org-uuid")
users, _, _ := scoped.Users.List(ctx, nil, nil)
```

## Pagination

```go
// Manual pagination
users, resp, _ := client.Users.List(ctx, nil, &jumpserver.ListOptions{
    Limit:  20,
    Offset: 0,
    Search: "admin",
})
if resp.HasNextPage() {
    // fetch next page...
}
// Auto-iterate all pages
var all []model.User
jumpserver.WalkPages(ctx, &jumpserver.ListOptions{Limit: 100}, 0,
    func(ctx context.Context, opts *jumpserver.ListOptions) (*jumpserver.Response, error) {
        users, resp, err := client.Users.List(ctx, nil, opts)
        if err != nil { return resp, err }
        all = append(all, users...)
        return resp, nil
    },
)
```

## Error Handling

```go
user, _, err := client.Users.Get(ctx, id)
if err != nil {
    if jumpserver.IsNotFound(err) {
        fmt.Println("user not found")
    }
    if jumpserver.IsUnauthorized(err) {
        fmt.Println("auth failed")
    }
    if jumpserver.IsRateLimited(err) {
        fmt.Println("rate limited")
    }
    var apiErr *jumpserver.APIError
    if errors.As(err, &apiErr) {
        fmt.Println(apiErr.StatusCode, apiErr.Message, string(apiErr.Body))
    }
}
```

## Retry

Default: 3 retries with exponential backoff and full jitter, respects `Retry-After` header:

```go
client := jumpserver.NewClient(
    jumpserver.WithBaseURL(url),
    jumpserver.WithRetry(5, 200*time.Millisecond, 30*time.Second),
    // ...
)
```

**Retried:**

- HTTP 408, 429, 500, 502, 503, 504
- Transient network errors (timeout, connection reset, temporary DNS failure)
  **Not retried:**
- `context.Canceled` / `context.DeadlineExceeded`
- TLS certificate errors
- 4xx client errors (except 408, 429)

## Services

| Service                                         | Client Field              | Methods                                                                                                                                                        | Description                                   |
| ----------------------------------------------- | ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------- |
| Auth                                            | `client.Auth`             | CreateToken / ConfirmLoginStatus / SelectMFA / CreateConnectionToken / CreateSuperConnectionToken / GetSuperConnectionTokenSecret / SSOLoginURL / GetClientURL | Login, MFA, connection tokens, SSO            |
| Users                                           | `client.Users`            | List / Get / Profile / Create / Update / Replace / Delete / Invite / ListGroups                                                                                | User CRUD + profile + group membership        |
| User Groups                                     | `client.UserGroups`       | List / Get / Create / Update / Delete / BindUsers / ListUsers                                                                                                  | Group CRUD + member management                |
| Roles                                           | `client.Roles`            | List(scope) / Get(scope)                                                                                                                                       | RBAC role queries by scope                    |
| Assets                                          | `client.Assets`           | List / Get / Delete / PermUsers                                                                                                                                | Generic asset operations + permitted users    |
| Hosts / Devices / DBs / Webs / Clouds / Customs | `client.Hosts` etc.       | List / Get / Create / Update / Replace / Delete                                                                                                                | Category-specific asset CRUD                  |
| Nodes                                           | `client.Nodes`            | List / Get / Create / Update / Delete / ChildrenTree / CreateChild                                                                                             | Asset tree node CRUD + children               |
| Platforms                                       | `client.Platforms`        | List / Get                                                                                                                                                     | Platform template queries (read-only)         |
| Zones                                           | `client.Zones`            | List / Get / Create / Update / Delete                                                                                                                          | Network zone CRUD                             |
| Gateways                                        | `client.Gateways`         | List / Get / Create / Update / Delete                                                                                                                          | Gateway CRUD                                  |
| Labels                                          | `client.Labels`           | List / Get / Create / Update / Delete                                                                                                                          | Label CRUD                                    |
| Accounts                                        | `client.Accounts`         | List / Get / Create / Update / Delete / GetSecret / CreateBulk / CreateBulkByTemplate / Verify / CreateVerifyTask                                              | Account CRUD + secret + bulk + connectivity   |
| Account Templates                               | `client.AccountTemplates` | List / Get / Create / Update / Delete                                                                                                                          | Template CRUD                                 |
| Change Secrets                                  | `client.ChangeSecrets`    | List / Get / Create / Update / Delete / Execute                                                                                                                | Secret rotation CRUD + execution              |
| Account Backups                                 | `client.AccountBackups`   | List / Get / Create / Update / Delete / Execute                                                                                                                | Backup plan CRUD + execution                  |
| Organizations                                   | `client.Orgs`             | List / Get / Create / Update / Delete                                                                                                                          | Organization CRUD                             |
| Permissions                                     | `client.Permissions`      | List / Get / Create / Update / Delete / GetSelfAssetAccounts / AddUsersRelations / AddUserGroupsRelations / AddAssetsRelations / AddNodesRelations             | Asset permission CRUD + batch relations       |
| Command Filters                                 | `client.CommandFilters`   | List / Get / Create / Update / Delete / ListGroups / GetGroup / CreateGroup / UpdateGroup / DeleteGroup                                                        | Command filter + command group CRUD           |
| Login ACLs                                      | `client.LoginACLs`        | List / Get                                                                                                                                                     | Login ACL queries (read-only)                 |
| Audits                                          | `client.Audits`           | ListSessions / GetSession / DownloadReplay / ListCommands / ListFTPLogs / ListLoginLogs / ListOperateLogs                                                      | Sessions, commands, FTP, login & operate logs |
| Terminal                                        | `client.Terminal`         | Register / Config / Heartbeat / ConnectMethods / GetTask                                                                                                       | Terminal component registration & config      |
| Tickets                                         | `client.Tickets`          | List / Get / Create / Approve / ListFlows / UpdateFlow                                                                                                         | Tickets + flow management                     |
| Settings                                        | `client.Settings`         | Public / List                                                                                                                                                  | System settings                               |
| Self                                            | `client.Self`             | ListAssets / GetAsset                                                                                                                                          | Self-service permitted assets                 |
| Ops                                             | `client.Ops`              | CreateJob / GetJobResult                                                                                                                                       | Ops job execution                             |
| Xpack                                           | `client.Xpack`            | License                                                                                                                                                        | Enterprise license info                       |

```
jumpserver-sdk-go/
├── client.go              # Client, HTTPClient interface
├── auth.go                # Authenticator implementations
├── options.go             # Functional options
├── errors.go              # APIError, error helpers
├── pagination.go          # ListOptions, Response, WalkPages
├── version.go             # SDK version
├── client_test.go         # Unit tests
├── Makefile               # Build/test/coverage commands
│
├── internal/core/         # Shared types (HTTPClient interface)
├── internal/util/      # Internal utilities
├── model/                 # Data models (pure type definitions)
│
├── auth/                  # Authentication service
├── users/                 # Users & groups (users.go, groups.go)
├── rbac/                  # Roles
├── assets/                # Assets/nodes/platforms/zones/gateways (7 files)
├── accounts/              # Accounts/templates/backup/change-secret (4 files)
├── orgs/                  # Organizations
├── perms/                 # Permissions (+ self assets)
├── ops/                   # Ops jobs
├── acls/                  # Command filters & login ACLs
├── audits/                # Audit logs (sessions, commands, ftplogs, logs)
├── terminal/              # Terminal
├── tickets/               # Tickets
├── settings/              # Settings
├── xpack/                 # Enterprise
├── labels/                # Labels
│
└── examples/
    ├── integration/       # Full CRUD integration test (200+ items)
    ├── list-users/
    ├── create-asset/
    └── connection-token/  # Connection token full flow
```

## Integration Test

Run the full CRUD test suite against a real JumpServer instance:

```bash
export JUMPSERVER_URL=https://your-jumpserver.example.com
export JUMPSERVER_KEY_ID=your-key-id
export JUMPSERVER_SECRET_ID=your-secret-id
make integration
# or directly
go run ./examples/integration
```

## Development

```bash
	make build       # Build all packages
	make test        # Run unit tests
	make vet         # Static analysis
	make fmt         # Format code (gofmt -w -s)
	make lint        # Static analysis (staticcheck)
	make all         # vet + test + build
	make precommit   # fmt + lint + vet + test + build
	make coverage    # Generate test coverage report
make clean       # Clean build artifacts
```

## Unit Test

```bash
go test ./...
```

## License

MIT — see [LICENSE](./LICENSE).
