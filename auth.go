package jumpserver

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Authenticator signs or annotates an outgoing *http.Request with
// credentials. Implementations must be safe for concurrent use.
type Authenticator interface {
	// Authenticate mutates req in place (typically by setting headers
	// or cookies). Returning a non-nil error aborts the request.
	Authenticate(req *http.Request) error
}

// SignatureAuth signs requests with HMAC-SHA256 using the JumpServer
// "Access Key" scheme, which is a profile of the IETF HTTP Signatures
// draft. The signed headers are "(request-target)" and "date".
type SignatureAuth struct {
	KeyID    string
	SecretID string
}

// Authenticate implements Authenticator.
func (a *SignatureAuth) Authenticate(req *http.Request) error {
	if a.KeyID == "" || a.SecretID == "" {
		return fmt.Errorf("jumpserver: SignatureAuth: KeyID and SecretID are required")
	}
	if req.Header.Get("Date") == "" {
		req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	}
	headers := []string{"(request-target)", "date"}
	lines := make([]string, 0, len(headers))
	for _, h := range headers {
		switch h {
		case "(request-target)":
			path := req.URL.RequestURI()
			lines = append(lines, fmt.Sprintf("(request-target): %s %s",
				strings.ToLower(req.Method), path))
		case "date":
			lines = append(lines, "date: "+req.Header.Get("Date"))
		}
	}
	sigString := strings.Join(lines, "\n")

	mac := hmac.New(sha256.New, []byte(a.SecretID))
	mac.Write([]byte(sigString))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	authz := fmt.Sprintf(
		`Signature keyId="%s",algorithm="hmac-sha256",headers="%s",signature="%s"`,
		a.KeyID, strings.Join(headers, " "), signature,
	)
	req.Header.Set("Authorization", authz)
	return nil
}

// BearerTokenAuth sets "Authorization: Bearer <token>".
type BearerTokenAuth struct {
	Token string
}

// Authenticate implements Authenticator.
func (a *BearerTokenAuth) Authenticate(req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+a.Token)
	return nil
}

// PrivateTokenAuth sets "Authorization: Token <token>".
type PrivateTokenAuth struct {
	Token string
}

// Authenticate implements Authenticator.
func (a *PrivateTokenAuth) Authenticate(req *http.Request) error {
	req.Header.Set("Authorization", "Token "+a.Token)
	return nil
}

// PasswordAuth authenticates by obtaining a Bearer token from the
// JumpServer authentication endpoint using username and password.
// The token is cached and automatically refreshed when it expires.
type PasswordAuth struct {
	// Username is the JumpServer username.
	Username string
	// Password is the JumpServer password.
	Password string

	mu          sync.Mutex
	baseURL     string
	httpClient  *http.Client
	token       string
	tokenExpiry time.Time
}

// init sets the base URL and HTTP client. Called by Client during
// initialization.
func (a *PasswordAuth) init(baseURL string, httpClient *http.Client) {
	a.baseURL = baseURL
	if httpClient != nil {
		a.httpClient = httpClient
	} else {
		a.httpClient = http.DefaultClient
	}
}

// Authenticate implements Authenticator. It obtains a Bearer token from
// the JumpServer authentication endpoint using the configured username
// and password. The token is cached until it expires.
func (a *PasswordAuth) Authenticate(req *http.Request) error {
	token, err := a.getToken()
	if err != nil {
		return fmt.Errorf("jumpserver: PasswordAuth: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

// getToken returns a valid token, refreshing it if necessary.
func (a *PasswordAuth) getToken() (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Return cached token if it's still valid (with 5-minute buffer).
	if a.token != "" && time.Now().Add(5*time.Minute).Before(a.tokenExpiry) {
		return a.token, nil
	}

	// Obtain a new token.
	token, expiry, err := a.fetchToken()
	if err != nil {
		return "", err
	}
	a.token = token
	a.tokenExpiry = expiry
	return token, nil
}

// fetchToken makes a POST request to the authentication endpoint to
// obtain a new token. Caller must hold a.mu.
func (a *PasswordAuth) fetchToken() (string, time.Time, error) {
	if a.baseURL == "" {
		return "", time.Time{}, fmt.Errorf("base URL not initialized")
	}

	baseURL := strings.TrimRight(a.baseURL, "/")
	url := baseURL + "/api/v1/authentication/auth/"

	body := map[string]string{
		"username": a.Username,
		"password": a.Password,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("marshal request body: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := a.httpClient.Do(httpReq)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("request token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", time.Time{}, fmt.Errorf("authentication failed: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Token       string `json:"token"`
		Keyword     string `json:"keyword"`
		DateExpired string `json:"date_expired"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", time.Time{}, fmt.Errorf("decode response: %w", err)
	}

	if result.Token == "" {
		return "", time.Time{}, fmt.Errorf("no token in response")
	}

	// Parse expiry time.
	expiry := time.Now().Add(24 * time.Hour) // Default to 24 hours if not specified.
	if result.DateExpired != "" {
		if t, err := time.Parse(time.RFC3339, result.DateExpired); err == nil {
			expiry = t
		}
	}

	return result.Token, expiry, nil
}
