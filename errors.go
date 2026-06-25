package jumpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// APIError represents a non-2xx response from the JumpServer API.
type APIError struct {
	StatusCode int
	Method     string
	URL        string
	Body       []byte
	Message    string
	Response   *http.Response
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("jumpserver: %s %s -> %d: %s", e.Method, e.URL, e.StatusCode, e.Message)
	}
	if len(e.Body) > 0 && len(e.Body) < 256 {
		return fmt.Sprintf("jumpserver: %s %s -> %d: %s", e.Method, e.URL, e.StatusCode, string(e.Body))
	}
	return fmt.Sprintf("jumpserver: %s %s -> %d", e.Method, e.URL, e.StatusCode)
}

// Is reports whether target is an *APIError with the same status code.
func (e *APIError) Is(target error) bool {
	var t *APIError
	if !errors.As(target, &t) {
		return false
	}
	if t.StatusCode != 0 && t.StatusCode != e.StatusCode {
		return false
	}
	return true
}

// Sentinel errors for common status classes.
var (
	ErrBadRequest   = &APIError{StatusCode: http.StatusBadRequest}
	ErrUnauthorized = &APIError{StatusCode: http.StatusUnauthorized}
	ErrForbidden    = &APIError{StatusCode: http.StatusForbidden}
	ErrNotFound     = &APIError{StatusCode: http.StatusNotFound}
	ErrConflict     = &APIError{StatusCode: http.StatusConflict}
	ErrRateLimited  = &APIError{StatusCode: http.StatusTooManyRequests}
	ErrServer       = &APIError{StatusCode: http.StatusInternalServerError}
)

func IsNotFound(err error) bool     { return errors.Is(err, ErrNotFound) }
func IsUnauthorized(err error) bool { return errors.Is(err, ErrUnauthorized) }
func IsForbidden(err error) bool    { return errors.Is(err, ErrForbidden) }
func IsRateLimited(err error) bool  { return errors.Is(err, ErrRateLimited) }

// extractAPIErrorMessage pulls a human-readable message from a JSON body.
// It handles Django REST Framework shapes:
//
//	{"detail": "..."}
//	{"message": "..."}
//	{"non_field_errors": ["..."]}
func extractAPIErrorMessage(body []byte) string {
	if len(body) == 0 {
		return ""
	}

	// Fast path: scan for "detail" or "message" string values.
	s := string(body)
	for _, key := range []string{`"detail":`, `"message":`, `"detail" :`, `"message" :`} {
		i := strings.Index(s, key)
		if i < 0 {
			continue
		}
		v := strings.TrimLeft(s[i+len(key):], " \t\n\r")
		if len(v) > 0 && v[0] == '"' {
			end := strings.IndexByte(v[1:], '"')
			if end >= 0 {
				return v[1 : end+1]
			}
		}
	}

	// Fallback: try JSON object with detail/message keys.
	var m map[string]any
	if json.Unmarshal(body, &m) == nil {
		for _, key := range []string{"detail", "message"} {
			if v, ok := m[key]; ok {
				return fmt.Sprint(v)
			}
		}
	}

	// Truncate long bodies.
	if len(body) > 200 {
		return string(body[:200]) + "..."
	}
	return s
}
