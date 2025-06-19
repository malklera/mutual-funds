package main

import (
	"log"
	"os"
)

const (
	fundsFile       = "funds.json"
	myFundsFile     = "myFunds.json"
	// WARN: change it back to 17
	hourCloseMarket = 20
)

func main() {
	switch len(os.Args) {
	case 1:
		menu()
	case 2:
		if os.Args[1] == "-u" {
			err := updateValues()
			if err != nil {
				log.Printf("Error with updateValues() : %v\n", err)
			}
		} else {
			log.Printf("Wrong argument: %s\n", os.Args[1])
		}
	default:
		log.Printf("Too many arguments.\n")
	}
}
