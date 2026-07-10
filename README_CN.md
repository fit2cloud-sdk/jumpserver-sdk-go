[![EN](https://img.shields.io/badge/EN-English-blue)](README.md)
[![CN](https://img.shields.io/badge/CN-中文-red)](README_CN.md)

# jumpserver-sdk-go

[JumpServer](https://www.jumpserver.org/) REST API 的 Go SDK，面向 **v4.10.x** 版本。

[English](README.md) | [中文文档](README_CN.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/fit2cloud-sdk/jumpserver-sdk-go.svg)](https://pkg.go.dev/github.com/fit2cloud-sdk/jumpserver-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/fit2cloud-sdk/jumpserver-sdk-go)](https://goreportcard.com/report/github.com/fit2cloud-sdk/jumpserver-sdk-go)

## 特性

- **完整 CRUD 覆盖** — 26 个服务模块，涵盖用户、资产、账号、权限、审计、工单、作业管理等全部核心功能
- **分类资产支持** — Hosts、Devices、Databases、Webs、Clouds、Customs 六大资产类别独立操作
- **多种认证方式** — AccessKey (HMAC-SHA256)、Bearer Token、Private Token、密码认证 (username/password)、自定义 Authenticator
- **组织作用域** — `WithOrgScope(id)` 切换组织上下文，无需重建 Client
- **自动分页** — `WalkPages()` 函数自动遍历所有分页
- **智能重试** — 指数退避 + 全抖动，仅重试瞬态错误（timeout、connection reset、429/5xx），永久错误不重试
- **零第三方依赖** — 纯标准库实现
- **Go 1.25** — 使用 `math/rand/v2`、`maps.Clone`、`for range int` 等新特性

## 环境要求

- Go 1.25+

## 安装

```bash
go get github.com/fit2cloud-sdk/jumpserver-sdk-go
```

## 快速开始

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

    // 列出用户
    users, _, err := client.Users.List(ctx, nil, &jumpserver.ListOptions{Limit: 20})
    if err != nil {
        log.Fatal(err)
    }
    for _, u := range users {
        fmt.Println(u.Username, u.Email)
    }

    // 按条件过滤
    users, _, _ = client.Users.List(ctx,
        map[string]string{"username": "admin"},
        &jumpserver.ListOptions{Limit: 10},
    )

    // 创建主机资产
    host, _, _ := client.Hosts.Create(ctx, &model.AssetRequest{
        Name:      "web-01",
        Address:   "192.168.1.10",
        Platform:  1, // Linux 平台 ID
        Protocols: []model.NamePort{{Name: "ssh", Port: 22}},
    })
    fmt.Println("Created:", host.ID)
}
```

## 认证方式

```go
// AccessKey HMAC-SHA256 签名（推荐，用于服务账号）
jumpserver.WithAccessKeyAuth(keyID, secretID)

// Bearer Token
jumpserver.WithBearerToken(token)

// Private Token (Authorization: Token <token>)
jumpserver.WithPrivateToken(token)

// 密码认证（用户名/密码，自动获取 Bearer token）
jumpserver.WithPasswordAuth(username, password)

// 自定义认证器
jumpserver.WithAuthenticator(myAuth)
```

## 组织作用域

JumpServer 的多端点通过组织路由。默认发送 `X-JMS-ORG: ROOT`。

```go
// 设置默认组织
client := jumpserver.NewClient(
    jumpserver.WithBaseURL(url),
    jumpserver.WithOrg("org-uuid"),
    // ...
)

// 派生一个作用域客户端（共享底层 HTTP 连接）
scoped := client.WithOrgScope("other-org-uuid")
users, _, _ := scoped.Users.List(ctx, nil, nil)
```

## 分页

```go
// 手动分页
users, resp, _ := client.Users.List(ctx, nil, &jumpserver.ListOptions{
    Limit:  20,
    Offset: 0,
    Search: "admin",
})
if resp.HasNextPage() {
    // 获取下一页...
}

// 自动遍历所有页
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

## 错误处理

```go
user, _, err := client.Users.Get(ctx, id)
if err != nil {
    // 按状态码判断
    if jumpserver.IsNotFound(err) {
        fmt.Println("user not found")
    }
    if jumpserver.IsUnauthorized(err) {
        fmt.Println("auth failed")
    }
    if jumpserver.IsRateLimited(err) {
        fmt.Println("rate limited")
    }

    // 获取详细信息
    var apiErr *jumpserver.APIError
    if errors.As(err, &apiErr) {
        fmt.Println(apiErr.StatusCode, apiErr.Message, string(apiErr.Body))
    }
}
```

## 重试

默认开启 3 次重试，指数退避 + 全抖动，遵守 `Retry-After` 响应头：

```go
client := jumpserver.NewClient(
    jumpserver.WithBaseURL(url),
    jumpserver.WithRetry(5, 200*time.Millisecond, 30*time.Second),
    // ...
)
```

**重试条件**：

- HTTP 408、429、500、502、503、504
- 瞬态网络错误（timeout、connection reset、DNS 临时故障）

**不重试**：

- `context.Canceled` / `context.DeadlineExceeded`
- TLS 证书错误
- 4xx 客户端错误（除 408、429）

## 服务列表

| 服务                                         | 字段                      | 方法                                                                                                                                                           | 说明                             |
| -------------------------------------------- | ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------- |
| 认证                                         | `client.Auth`             | CreateToken / ConfirmLoginStatus / SelectMFA / CreateConnectionToken / CreateSuperConnectionToken / GetSuperConnectionTokenSecret / SSOLoginURL / GetClientURL | 登录、MFA、连接令牌、SSO         |
| 用户                                         | `client.Users`            | List / Get / Profile / Create / Update / Replace / Delete / Invite / ListGroups                                                                                | 用户 CRUD + 个人信息 + 组成员    |
| 用户组                                       | `client.UserGroups`       | List / Get / Create / Update / Delete / BindUsers / ListUsers                                                                                                  | 用户组 CRUD + 成员管理           |
| 角色                                         | `client.Roles`            | List(scope) / Get(scope)                                                                                                                                       | 按作用域查询 RBAC 角色           |
| 资产                                         | `client.Assets`           | List / Get / Delete / PermUsers                                                                                                                                | 通用资产操作 + 授权用户          |
| 主机 / 网络设备 / 数据库 / Web / 云 / 自定义 | `client.Hosts` 等         | List / Get / Create / Update / Replace / Delete                                                                                                                | 分类资产 CRUD                    |
| 节点                                         | `client.Nodes`            | List / Get / Create / Update / Delete / ChildrenTree / CreateChild                                                                                             | 资产树节点 CRUD + 子节点         |
| 平台                                         | `client.Platforms`        | List / Get                                                                                                                                                     | 平台模板查询（只读）             |
| 网域                                         | `client.Zones`            | List / Get / Create / Update / Delete                                                                                                                          | 网域 CRUD                        |
| 网关                                         | `client.Gateways`         | List / Get / Create / Update / Delete                                                                                                                          | 网关 CRUD                        |
| 标签                                         | `client.Labels`           | List / Get / Create / Update / Delete                                                                                                                          | 标签 CRUD                        |
| 账号                                         | `client.Accounts`         | List / Get / Create / Update / Delete / GetSecret / CreateBulk / CreateBulkByTemplate / Verify / CreateVerifyTask                                              | 账号 CRUD + 密钥 + 批量 + 连通性 |
| 账号模板                                     | `client.AccountTemplates` | List / Get / Create / Update / Delete                                                                                                                          | 模板 CRUD                        |
| 改密自动化                                   | `client.ChangeSecrets`    | List / Get / Create / Update / Delete / Execute                                                                                                                | 改密策略 CRUD + 执行             |
| 账号备份                                     | `client.AccountBackups`   | List / Get / Create / Update / Delete / Execute                                                                                                                | 备份计划 CRUD + 执行             |
| 组织                                         | `client.Orgs`             | List / Get / Create / Update / Delete                                                                                                                          | 组织 CRUD                        |
| 权限                                         | `client.Permissions`      | List / Get / Create / Update / Delete / GetSelfAssetAccounts / AddUsersRelations / AddUserGroupsRelations / AddAssetsRelations / AddNodesRelations             | 资产授权 CRUD + 批量关系         |
| 命令过滤                                     | `client.CommandFilters`   | List / Get / Create / Update / Delete / ListGroups / GetGroup / CreateGroup / UpdateGroup / DeleteGroup                                                        | 命令过滤 + 命令组 CRUD           |
| 登录 ACL                                     | `client.LoginACLs`        | List / Get                                                                                                                                                     | 登录 ACL 查询（只读）            |
| 审计                                         | `client.Audits`           | ListSessions / GetSession / DownloadReplay / ListCommands / ListFTPLogs / ListLoginLogs / ListOperateLogs                                                      | 会话、命令、FTP、登录、操作日志  |
| 终端                                         | `client.Terminal`         | Register / Config / Heartbeat / ConnectMethods / GetTask                                                                                                       | 终端注册与配置                   |
| 工单                                         | `client.Tickets`          | List / Get / Create / Approve / ListFlows / UpdateFlow                                                                                                         | 工单 + 流程管理                  |
| 设置                                         | `client.Settings`         | Public / List                                                                                                                                                  | 系统设置                         |
| 我的资产                                     | `client.Self`             | ListAssets / GetAsset                                                                                                                                          | 当前用户授权的资产               |
| 作业                                         | `client.Ops`              | CreateJob / GetJobResult                                                                                                                                       | 作业执行                         |
| 企业版                                       | `client.Xpack`            | License                                                                                                                                                        | 企业版 License 信息              |

## 包结构

```
jumpserver-sdk-go/
├── client.go              # Client、HTTPClient 接口
├── auth.go                # 认证器实现
├── options.go             # 函数式配置
├── errors.go              # APIError、错误判断辅助函数
├── pagination.go          # ListOptions、Response、WalkPages
├── version.go             # SDK 版本号
├── client_test.go         # 单元测试
├── Makefile               # 构建/测试/覆盖率等常用命令
│
├── internal/core/         # 共享类型（HTTPClient 接口）
├── internal/util/      # 内部工具函数
├── model/                 # 数据模型（纯类型定义）
│
├── auth/                  # 认证服务
├── users/                 # 用户 & 用户组（users.go, groups.go）
├── rbac/                  # 角色
├── assets/                # 资产/节点/平台/网域/网关（7 个文件）
├── accounts/              # 账号/模板/改密/备份（4 个文件）
├── orgs/                  # 组织
├── perms/                 # 权限（+ 我的资产）
├── ops/                   # 作业
├── acls/                  # 命令过滤 & 登录 ACL
├── audits/                # 审计日志（sessions, commands, ftplogs, logs）
├── terminal/              # 终端
├── tickets/               # 工单
├── settings/              # 设置
├── xpack/                 # 企业版
├── labels/                # 标签
│
└── examples/
    ├── integration/       # 完整 CRUD 集成测试（200+ 项）
    ├── list-users/
    ├── create-asset/
    └── connection-token/  # 连接令牌完整流程
```

## 集成测试

使用真实 JumpServer 实例运行全量 CRUD 测试：

```bash
export JUMPSERVER_URL=https://your-jumpserver.example.com
export JUMPSERVER_KEY_ID=your-key-id
export JUMPSERVER_SECRET_ID=your-secret-id

make integration
# 或直接运行
go run ./examples/integration
```

## 开发

```bash
	make build       # 编译所有包
	make test        # 运行单元测试
	make vet         # 静态检查
	make fmt         # 格式化代码 (gofmt -w -s)
	make lint        # 静态分析 (staticcheck)
	make all         # vet + test + build
	make precommit   # fmt + lint + vet + test + build
	make coverage    # 生成测试覆盖率报告
make clean       # 清理编译产物
```

## 单元测试

```bash
go test ./...
```

## 许可证

MIT — see [LICENSE](./LICENSE).
