// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/KingJorjai/BONK/pkg/cli"
	"github.com/KingJorjai/BONK/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	checkUpdates bool
	leaderboard  bool
	version      bool
)

var rootCmd = &cobra.Command{
	Use:   "BONK [name]",
	Short: "BONK is a tool for bonking people",
	Long:  `BONK is a CLI tool for bonking people and showing leaderboards.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// If the --version flag is set, show the version
		if version {
			fmt.Printf("BONK v%s\n", utils.AppVersion)
			return
		}
		// If the --leaderboard flag is set, show the leaderboard
		if leaderboard {
			cli.ShowLeaderboard()
			return
		}

		if checkUpdates {
			cli.CheckForUpdates()
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
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.DisableDescriptions = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	rootCmd.DisableAutoGenTag = true

	rootCmd.Flags().BoolVarP(&leaderboard, "leaderboard", "l", false, "show the most bonked in a leaderboard")
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show the current version")
	rootCmd.Flags().BoolVarP(&checkUpdates, "check-updates", "u", false, "checks for new updates")

	rootCmd.MarkFlagsMutuallyExclusive("leaderboard", "version", "check-updates")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
