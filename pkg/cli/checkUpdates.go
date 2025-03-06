package cli

import (
	"fmt"

	"github.com/KingJorjai/BONK/pkg/api"
	"github.com/KingJorjai/BONK/pkg/utils"
)

func CheckForUpdates() {
	update, version, err := api.CheckForUpdate(utils.AppVersion)

	if err != nil {
		fmt.Printf("Error checking for updates: %v\n", err)
		return
	}

	if update {
		fmt.Printf("A new version is available: v%s\n", version)
	} else {
		fmt.Println("You are already using the latest version.")
	}
}
