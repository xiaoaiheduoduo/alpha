package ahttp

import (
	"net/http"
	"strings"
)

var textContentTypes = []string{"application/json", "text/html", "application/xml", "text/xml", "text/plain", "application/x-www-form-urlencoded", "application/x-yaml"}

func IsTextContentType(contentType string) bool {
	ct := strings.ToLower(contentType)
	for _, typ := range textContentTypes {
		if ct == typ {
			return true
		}
	}

	return false
}

func FilterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func GetContentType(headers http.Header) string {
	if len(headers) < 1 {
		return ""
	}
	v, ok := headers["Content-Type"]
	if !ok {
		return ""
	}
	if len(v) > 0 {
		return FilterFlags(v[0])
	}
	return ""
}
