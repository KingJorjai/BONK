package main

import (
	"fmt"
	"os"

	"github.com/KingJorjai/BONK/pkg/cli"
)

func main() {
	// Check if there's exactly one command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <name>")
		os.Exit(2)
	}

	// Store the argument
	name := os.Args[1]

	cli.Bonk(name)
}
