package main

import (
	"fmt"
	"os"

	"github.com/KingJorjai/BONK/pkg/cli"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <name>")
		os.Exit(2)
	}
	name := os.Args[1]

	cli.Bonk(name)
}
