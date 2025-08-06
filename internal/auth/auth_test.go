package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name          string
		header        http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "Success",
			header:        http.Header{"Authorization": []string{"ApiKey someValidKey123"}},
			expectedKey:   "someValidKey12",
			expectedError: nil,
		},
		{
			name:          "No Auth Header",
			header:        http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed Auth",
			header:        http.Header{"Authorization": []string{"someValidKey123"}},
			expectedKey:   "",
			expectedError: ErrMalformedAuth,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotKey, goterr := GetAPIKey(tc.header)

			if gotKey != tc.expectedKey {
				t.Errorf("got: %v want: %v", gotKey, tc.expectedKey)
			}

			if goterr != tc.expectedError {
				t.Errorf("got: %v want: %v", goterr, tc.expectedError)
			}
		})
	}
}
