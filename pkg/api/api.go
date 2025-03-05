package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Bonkee struct {
	Name      string    `json:"name"`
	Bonks     uint      `json:"bonks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
	resp, err := http.Get(fmt.Sprintf("%s/bonk/%s", os.Getenv("BONK_API_URL"), name))
	if err != nil {
		return 0, fmt.Errorf("error while calling the API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var bonkNumberJson Bonkee
	if err := json.NewDecoder(resp.Body).Decode(&bonkNumberJson); err != nil {
		return 0, fmt.Errorf("error while decoding the response: %w", err)
	}

	return int(bonkNumberJson.Bonks), nil
}
