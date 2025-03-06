// main.go
package main

import (
	"os"

	"github.com/KingJorjai/BONK/cmd"
)

const (
	AppVersion = "1.0.3"
)

func main() {
	// Set up environment variables
	os.Setenv("BONK_API_URL", "http://143.47.47.64:25572/api")

	// Execute the root command
	cmd.Execute()
}
