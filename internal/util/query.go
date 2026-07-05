package util

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// AppendQuery builds "path?key=val&..." from a map of params.
func AppendQuery(path string, params map[string]string) string {
	if len(params) == 0 {
		return path
	}
	q := url.Values{}
	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	encoded := q.Encode()
	if encoded == "" {
		return path
	}
	if strings.ContainsRune(path, '?') {
		return path + "&" + encoded
	}
	return path + "?" + encoded
}

// Itoa converts n to its string representation.
func Itoa(n int) string { return strconv.Itoa(n) }

// Spath formats a path template using fmt.Sprintf.
func Spath(tpl string, args ...any) string {
	return fmt.Sprintf(tpl, args...)
}
