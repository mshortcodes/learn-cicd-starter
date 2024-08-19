package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		input    http.Header
		expected string
	}{
		{
			input:    http.Header{"Authorization": []string{"ApiKey 12345"}},
			expected: "12345",
		},
		{
			input:    http.Header{"Authorization": []string{"ApiKey12345"}},
			expected: " ",
		},
	}

	for _, c := range cases {
		actual, _ := GetAPIKey(c.input)
		if actual != c.expected {
			t.Errorf("ApiKey doesn't match: %s vs %s", actual, c.expected)
		}
	}
}
