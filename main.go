package main

import (
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		// Show the menu here
		fmt.Println("The menu")
	case 2:
		if os.Args[1] == "-u" {
			createFundsFile(file)
			saveValues()
		} else {
			fmt.Println("Wrong argument:", os.Args[1])
		}
	default:
		fmt.Println("Too many arguments.")
	}
}
