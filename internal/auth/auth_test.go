package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			headers:       http.Header{"Authorization": []string{"ApiKey valid"}},
			expectedKey:   "valid",
			expectedError: nil,
		},
		{
			headers:       http.Header{"Authorization": []string{"invalid"}},
			expectedKey:   "",
			expectedError: ErrMalformedToken,
		},
		{
			headers:       http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey:   "",
			expectedError: ErrMalformedToken,
		},
	}

	for _, b := range tests {
		key, err := GetAPIKey(b.headers)
		if key != b.expectedKey || err != b.expectedError {
			t.Errorf("Actual:\n\tkey: %v, err: %v\nExpected:\n\tkey: %v, err: %v", key, err, b.expectedKey, b.expectedError)
		}
	}
}
