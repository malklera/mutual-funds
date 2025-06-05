package main

import (
	"fmt"
	"os"
)

const (
	fundsFile  = "funds.json"
	myFundsFile = "myFunds.json"
)

func main() {
	switch len(os.Args) {
	case 1:
		menu()
	case 2:
		if os.Args[1] == "-u" {
			createBaseFile(fundsFile, baseFundsJson)
			createBaseFile(myFundsFile, baseMyFundsJson)
			saveValues()
		} else {
			fmt.Println("Wrong argument:", os.Args[1])
		}
	default:
		fmt.Println("Too many arguments.")
	}
}
