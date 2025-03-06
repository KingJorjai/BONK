package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KingJorjai/BONK/pkg/cli"
)

func main() {
	// Set up environment variables
	os.Setenv("BONK_API_URL", "http://143.47.47.64:25572/api")

	// Parse command line flags
	leaderboardFlag := flag.Bool("leaderboard", false, "Show the most bonked in a leaderboard")
	flag.Parse()

	// Flag specific execution
	if *leaderboardFlag {
		cli.ShowLeaderboard()
		os.Exit(0)
	}

	// Default execution
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: " + os.Args[0] + " <name>")
		os.Exit(2)
	}
	name := args[0]
	cli.Bonk(name)

	os.Exit(0)
}
