package main

import (
	"fmt"
	"os"
)

const (
	fundsFile   = "funds.json"
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
			// TODO: Check if i already run this today after the close of market, if yes abort
			// otherwise continue, where do i check that?, should i get the hour and minute when
			// the data is save instead of only the yyyy/mm/dd ?
			// this is because some days i will run the program at morning and maybe or not
			// at night

			updateValues()
		} else {
			fmt.Println("Wrong argument:", os.Args[1])
		}
	default:
		fmt.Println("Too many arguments.")
	}
}
