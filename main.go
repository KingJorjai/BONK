package main

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
	// The URL of the API
	API_URL       = "https://api.counterapi.dev/v1"
	API_NAMESPACE = "bonk.jorjai.net"
	ASCII_ART     = `⠀⠀⠀⠀⠀⠀⢀⣁⣤⣶⣶⡒⠒⠲⠾⣭⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣿⡀⣸⠟⠛⠃⠀⣀⣀⠈⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⡠⠂⢠⠏⠀⠉⠀⠀⠀⠰⣿⠟⠀⠙⢧⡀⠀⠀⠀⠀⠀⠀⢀⠀⠀⢀⢀⡀⣼⣧⡾⠃⠀⠀⠀⠀⠀
⢀⠔⠀⣠⠔⠁⠀⠀⠀⠀⠀⠀⠀⠰⢄⡠⣶⢾⣽⡆⠀⠀⠀⠀⠄⢡⡀⢰⣾⣿⡀⠈⠵⠟⠛⠀⠀⠀⠀⠀
⠀⣠⠊⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⡟⠋⠉⠀⠀⠀⠀⣰⢦⣼⡷⣼⡏⢯⢉⣡⠖⠋⣩⡇⠀⠀⠀⠀
⣰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⠀⢺⣿⡄⣿⡄⣿⢿⠈⢁⡴⠋⠀⢀⣴⣋⡀⠀⠀⠀⠀
⡇⠀⠀⢰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠒⢛⣡⣸⡏⢹⡟⠻⠏⢀⡴⠋⠀⣠⣖⠻⠿⠿⣤⡀⠀⠀⠀
⡇⠀⠀⠈⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡟⠀⠛⢡⡞⠻⠟⠁⢀⡴⠋⢀⣤⣞⣛⣻⡆⠀⠀⠉⢇⠀⠀⠀
⣇⠀⠀⠀⠈⢦⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⠀⠠⠏⠀⠀⢀⠴⠋⣀⠴⣿⠛⠛⠁⠈⠁⠀⠀⠀⠈⢧⠀⡄
⠸⡄⠀⠀⠀⠀⡇⠀⠀⢰⠃⠀⠈⣇⠀⠸⣦⡀⠀⠀⢀⡔⠁⣠⠞⠁⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡀⠀
⠀⠙⣄⠀⠀⠀⣿⠀⠀⢸⠒⠒⠒⠻⡀⠀⣷⠬⣉⡶⠋⣠⠞⠁⠀⠀⠀⡇⠀⠀⠀⡀⠀⢠⠀⠀⠀⠘⡇⠀
⠀⠀⠈⠑⠦⠤⣽⣄⠀⢸⠤⠤⠤⠤⢷⡀⠸⣷⠋⣠⢾⡁⠀⠀⠀⠀⠀⡇⢠⠇⠀⢹⠀⢸⠃⠀⠀⣸⠃⠀
⠀⠀⠀⠀⠀⠀⠀⢹⠀⢸⠀⠀⠀⠀⠀⢈⠦⣀⣙⣻⡞⠃⠀⠀⠀⢀⡼⢡⠧⠤⠤⢸⠀⣾⠤⠤⠚⠁⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢸⡀⠸⡄⠀⠀⠀⠀⣧⠴⠃⠉⠉⠁⠀⠀⠰⣾⡭⠔⠁⠀⠀⠀⡜⠀⡇⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠳⢤⣼⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠂⠄⠀⠀⠀⠀⠀⢰⣥⣴⠃⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⠤⠐⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`
)

func main() {
	// Check if there's exactly one command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: bonk <name>")
		os.Exit(2)
	}

	// Store the argument
	name := os.Args[1]

	// Use the API to increment the counter and get the new value
	bonkNumber := bonk(name)

	fmt.Println(ASCII_ART)
	fmt.Printf("You bonked %s.\n", name)
	fmt.Printf("%s was bonked %d times.\n", name, bonkNumber)
	os.Exit(0)
}

func bonk(name string) int {
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
