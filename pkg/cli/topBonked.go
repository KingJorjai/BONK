package cli

import (
	"fmt"

	"github.com/KingJorjai/BONK/pkg/api"
)

func TopBonked() {
	// Retrieve the top 10 from the API
	resp, err := api.GetTopBonked()
	if err != nil {
		fmt.Println("Unable to find the most bonked... Try later.")
		return
	}

	// Format to a table and print
	fmt.Println("+-----+--------------------------+-------+")
	fmt.Println("| NO. |           NAME           | BONKS |")
	fmt.Println("+-----+--------------------------+-------+")
	for i, bonk := range resp {
		name := bonk.Name
		if len(name) > 21 {
			name = name[:21] + "..."
		}
		fmt.Printf("| %3d | %24s | %-5d |\n", i+1, name, bonk.Bonks)
	}
	fmt.Println("+-----+--------------------------+-------+")
}
