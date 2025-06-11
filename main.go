package main

import (
	"log"
	"os"
)

const (
	fundsFile       = "funds.json"
	myFundsFile     = "myFunds.json"
	hourCloseMarket = 17
)

func main() {
	switch len(os.Args) {
	case 1:
		menu()
	case 2:
		if os.Args[1] == "-u" {
			updateValues()
		} else {
			log.Fatalf("Wrong argument: %s", os.Args[1])
		}
	default:
		log.Fatalf("Too many arguments.")
	}
}
