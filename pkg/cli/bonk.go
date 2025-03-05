package cli

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/KingJorjai/BONK/pkg/api"
)

const (
	ASCII_BONK = `⠀⠀⠀⠀⠀⠀⢀⣁⣤⣶⣶⡒⠒⠲⠾⣭⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
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
	ASCII_NO_BONK = `⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣶⡒⠒⠲⠾⣭⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣿⡀⣸⠟⠛⠃⠀⣀⣀⠈⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢠⠏⠀⠉⠀⠀⠀⠰⣿⠟⠀⠙⢧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣠⠔⠁⠀⠀⠀⠀⠀⠀⠀⠰⢄⡠⣶⢾⣽⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⣠⠊⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⡟⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⠖⠋⣩⡇⠀⠀⠀⠀
⣰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠋⠀⢀⣴⠋⠀⠀⠀⠀⠀
⡇⠀⠀⢰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠋⠀⣠⣖⠓⠀⠀⠀⠀⠀⠀⠀
⡇⠀⠀⠈⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡟⠀⠀⠀⠀⠀⠀⠀⢀⡴⠋⢀⣤⡞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣇⠀⠀⠀⠈⢦⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⠀⠀⠀⠀⠀⢀⠴⠋⣀⠴⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠸⡄⠀⠀⠀⠀⡇⠀⠀⢰⠃⠀⠈⣇⠀⠸⣦⡀⠀⠀⢀⡔⠁⣠⠞⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠙⣄⠀⠀⠀⣿⠀⠀⢸⠒⠒⠒⠻⡀⠀⣷⠬⣉⡶⠋⣠⠞⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠈⠑⠦⠤⣽⣄⠀⢸⠤⠤⠤⠤⢷⡀⠸⣷⠋⣠⢾⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢹⠀⢸⠀⠀⠀⠀⠀⢈⠦⣀⣙⣻⡞⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢸⡀⠸⡄⠀⠀⠀⠀⣧⠴⠃⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠳⢤⣼⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`
)

// Bonk performs a "bonk" operation on the specified entity.
//
// It validates that the provided name contains only letters, numbers, and spaces
// and is between 1-255 characters long. If validation fails, it prints an error
// message and exits with code 1.
//
// If validation passes, it increments a counter for this name using api.CountUp,
// displays ASCII art, outputs bonking messages including the total number of times
// the entity has been bonked, and exits with code 0.
//
// Parameters:
//   - name: The name of the entity to bonk (must be 1-255 alphanumeric characters)
//
// Side effects:
//   - Prints to stdout
//   - Terminates the program with exit code 1 on validation failure
func Bonk(name string) {
	// Remove all spaces longer than one character and trim
	name = strings.Join(strings.Fields(name), " ")

	// Validate that name only contains alphanumeric characters and is between 1-255 characters
	match, _ := regexp.MatchString("^[\\p{L}\\p{N} ]{1,255}$", name)
	if !match {
		fmt.Println(ASCII_NO_BONK)
		fmt.Println("BONK DENIED! That name is too sus. Use only letters and numbers (1-255 characters). No bonking the void!")
		os.Exit(1)
	}

	bonkNumber, err := api.CountUp(name)

	if err != nil {
		fmt.Println(ASCII_NO_BONK)
		fmt.Printf("No BONK :(")
		os.Exit(1)

	}

	fmt.Println(ASCII_BONK)
	fmt.Printf("You gave %s a mighty BONK!\n", name)
	fmt.Printf("%s has been bonked %d time%s! BONK BONK!\n", name, bonkNumber, pluralize(bonkNumber))
}

// pluralize returns "s" if the count is not equal to 1, otherwise it returns an empty string.
// This is useful for pluralizing words based on the count.
func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}
