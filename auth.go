package jumpserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
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

// BasicAuth sets HTTP Basic credentials on the request.
type BasicAuth struct {
	Username string
	Password string
}

// Authenticate implements Authenticator.
func (a *BasicAuth) Authenticate(req *http.Request) error {
	req.SetBasicAuth(a.Username, a.Password)
	return nil
}
