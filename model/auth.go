package model

// TokenRequest is the body for authentication token creation
// (username+password login).
type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	MFAType  string `json:"mfa_type,omitempty"`
	MFACode  string `json:"mfa_code,omitempty"`
}

// Token is the result of a successful authentication.
type Token struct {
	Token       string `json:"token"`
	Keyword     string `json:"keyword"`
	DateExpired string `json:"date_expired"`
	User        any    `json:"user"`
}

// ConnectionTokenRequest is the body for creating a connection token.
type ConnectionTokenRequest struct {
	User           string `json:"user,omitempty"`
	Asset          string `json:"asset"`
	Account        string `json:"account"`
	Protocol       string `json:"protocol"`
	ConnectMethod  string `json:"connect_method"`
	InputUsername  string `json:"input_username,omitempty"`
	InputSecret    string `json:"input_secret,omitempty"`
	ConnectOptions any    `json:"connect_options,omitempty"`
}

// ConnectionToken is the result of a connection token creation.
type ConnectionToken struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	User        any    `json:"user"`
	Asset       any    `json:"asset"`
	Account     any    `json:"account"`
	Protocol    string `json:"protocol"`
	ExpireAt    int64  `json:"expire_at"`
	IsActive    bool   `json:"is_active"`
	OrgName     string `json:"org_name"`
	DateCreated string `json:"date_created"`
}

// SSOLoginRequest is the body for requesting an SSO login URL.
type SSOLoginRequest struct {
	Username string `json:"username"`
	Next     string `json:"next,omitempty"`
}
