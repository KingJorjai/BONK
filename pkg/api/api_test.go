package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KingJorjai/BONK/pkg/api"
)

func TestGetLatestVersion(t *testing.T) {
	tests := []struct {
		name            string
		responseBody    string
		responseStatus  int
		expectedVersion string
		expectError     bool
	}{
		{"ValidResponse", `{"tag_name": "v1.1.0"}`, http.StatusOK, "1.1.0", false},
		{"InvalidJSON", `{"tag_name": "v1.1.0"`, http.StatusOK, "", true},
		{"Non200Status", ``, http.StatusNotFound, "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.responseStatus)
				w.Header().Set("Content-Type", "application/json")
				_, err := w.Write([]byte(test.responseBody)) // Check for errors
				if err != nil {
					t.Fatalf("failed to write response: %v", err)
				}
			}))
			defer server.Close()

			apiClient := api.GithubApiClient{Endpoint: server.URL}
			version, err := apiClient.GetLatestVersion()

			if (err != nil) != test.expectError {
				t.Errorf("expected error: %v, got: %v", test.expectError, err)
			}
			if version != test.expectedVersion {
				t.Errorf("expected version: %s, got: %s", test.expectedVersion, version)
			}
		})
	}
}

func TestCheckForUpdate(t *testing.T) {
	tests := []struct {
		name           string
		currentVersion string
		latestVersion  string
		expectedUpdate string
		expectError    bool
	}{
		{"UpdateAvailable", "1.0.0", "1.1.0", "1.1.0", false},
		{"NoUpdateAvailable", "1.1.0", "1.1.0", "", true},
		{"InvalidCurrentVersion", "invalid", "1.1.0", "1.1.0", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"tag_name": "v` + test.latestVersion + `"}`))
			}))
			defer server.Close()

			apiClient := api.GithubApiClient{Endpoint: server.URL}
			update, err := apiClient.CheckForUpdate(test.currentVersion)

			if (err != nil) != test.expectError {
				t.Errorf("expected error: %v, got: %v", test.expectError, err)
			}
			if update != test.expectedUpdate {
				t.Errorf("expected update: %s, got: %s", test.expectedUpdate, update)
			}
		})
	}
}
