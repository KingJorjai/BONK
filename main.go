package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KingJorjai/BONK/pkg/cli"
)

func main() {
	// Set up environment variables
	os.Setenv("BONK_API_URL", "https://api.counterapi.dev/v1")
	os.Setenv("BONK_API_NAMESPACE", "bonk.jorjai.net")

	// Parse command line flags
	topBonked := flag.Bool("top", false, "top")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: " + os.Args[0] + " <name>")
		os.Exit(2)
	}
	name := args[0]
	cli.Bonk(name)

	if *topBonked {
		// TODO Implement
	}

	os.Exit(0)
}
