// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/KingJorjai/BONK/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	leaderboard bool
)

var rootCmd = &cobra.Command{
	Use:   "BONK [name]",
	Short: "BONK is a tool for bonking people",
	Long:  `BONK is a CLI tool for bonking people and showing leaderboards.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// If the --leaderboard flag is set, show the leaderboard
		if leaderboard {
			cli.ShowLeaderboard()
			return
		}

		// If no name is specified, print usage
		if len(args) == 0 {
			cmd.Usage()
			os.Exit(2)
		}

		// Bonk the specified name
		name := args[0]
		cli.Bonk(name)
	},
}

func init() {
	// Define the --leaderboard flag
	rootCmd.Flags().BoolVarP(&leaderboard, "leaderboard", "l", false, "Show the most bonked in a leaderboard")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
