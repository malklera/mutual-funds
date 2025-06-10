package main

import (
	"log"
	"os"
	"time"
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

			stat, err := os.Stat(fundsFile)
			if err != nil {
				log.Fatalf("Error reading the stats of: %s : %v", fundsFile, err)
			}
			fileModTime := stat.ModTime()
			now := time.Now()
			fileModDay := time.Date(fileModTime.Year(), fileModTime.Month(),
				fileModTime.Day(), 0, 0, 0, 0, fileModTime.Location())
			today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

			// NOTE: for the future i will like to work with RFC3339 format, i just do
			// not like the way i would do it now

			if fileModDay.Before(today) {
				updateValues()
			} else {
				closeMarket := time.Date(now.Year(), now.Month(), now.Day(), hourCloseMarket,
					0, 0, 0, now.Location())

				if fileModTime.Before(closeMarket) {
					updateValues()
				} else {
					return
				}
			}
		} else {
			log.Fatalf("Wrong argument: %s", os.Args[1])
		}
	default:
		log.Fatalf("Too many arguments.")
	}
}
