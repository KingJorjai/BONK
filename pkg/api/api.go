package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
//
// Note: This function will terminate the program with os.Exit(1) if any errors
// occur during the API request or response parsing.
func CountUp(name string) int {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/up", API_URL, API_NAMESPACE, name))
	if err != nil {
		fmt.Println("Error while calling the API")
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var bonkNumberJson CountApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&bonkNumberJson); err != nil {
		fmt.Println("Error while decoding the response")
		os.Exit(1)
	}

	return bonkNumberJson.Count
}
