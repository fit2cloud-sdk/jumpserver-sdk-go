# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/).

## [0.1.0] - 2026-06-21

### Added
- Initial public release of `jumpserver-sdk-go`.
- Root `Client` with typed services: Auth, Users, UserGroups, Roles,
  Assets, Hosts, Devices, Databases, Webs, Clouds, Customs, Nodes,
  Platforms, Domains, Zones, Gateways, Labels, Accounts, Organizations,
  Permissions, CommandFilters, LoginACLs, Audits, Terminal, Tickets,
  Settings, Xpack.
- Pluggable authentication: `SignatureAuth` (HMAC-SHA256), `BearerTokenAuth`,
  `PrivateTokenAuth`, `PasswordAuth` (username/password), and custom `Authenticator`.
- Typed pagination via `ListOptions` and `Response`; `Client.WalkPages`
  auto-pagination helper.
- Typed `APIError` with `IsNotFound`, `IsUnauthorized`, `IsForbidden`,
  `IsRateLimited` helpers.
- Retry with exponential backoff + jitter on 408/429/5xx and transient
  network errors; honours `Retry-After`.
- `Client.WithOrgScope(id)` for per-call organization overrides.
- `Client.DoRaw` for streaming binary responses (e.g. session replays).
- Zero third-party runtime dependencies.
