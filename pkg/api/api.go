package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Namespace struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CountApiResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	NamespaceID int       `json:"namespace_id"`
	Namespace   Namespace `json:"namespace"`
}

const (
	API_URL       = "https://api.counterapi.dev/v1"
	API_NAMESPACE = "bonk.jorjai.net"
)

// CountUp increments a counter for the given name by sending a request to the API.
// It retrieves the current count value associated with the name parameter.
//
// Parameters:
//   - name: The identifier for the counter to increment
//
// Returns:
//   - The updated count value after incrementing
//   - An error if any occurs during the API request or response parsing.
func CountUp(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/up", API_URL, API_NAMESPACE, name))
	if err != nil {
		return 0, fmt.Errorf("error while calling the API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var bonkNumberJson CountApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&bonkNumberJson); err != nil {
		return 0, fmt.Errorf("error while decoding the response: %w", err)
	}

	return bonkNumberJson.Count, nil
}
