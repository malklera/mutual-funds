package main

import (
	"fmt"
	"os"
)

func main() {
	switch {
	case len(os.Args) == 1:
		// Show the menu here
	case len(os.Args) == 2:
		if os.Args[1] == "-u" {
			// Run saveValues()
			fmt.Println("The argument is:", os.Args[1])
		} else {
			fmt.Println("Wrong argument:", os.Args[1])
		}
	default:
		fmt.Println("Too many arguments.")
	}
}
