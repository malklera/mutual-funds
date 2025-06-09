package main

import (
	"log"
	"os"
	"time"
)

const (
	fundsFile       = "funds.json"
	myFundsFile     = "myFunds.json"
	hourCloseMarket = "17:00:00"
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

			today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0,
				now.Location())

			if fileModDay.Before(today) {
				createBaseFile(fundsFile, baseFundsJson)
				createBaseFile(myFundsFile, baseMyFundsJson)
				updateValues()
			} else {
				parsedHourCloseMarket, err := time.Parse(time.TimeOnly, hourCloseMarket)
				if err != nil {
					log.Fatalf("Error parsing %s : %v", hourCloseMarket, err)
				}

				marketCloseToday := time.Date(now.Year(), now.Month(), now.Day(),
					parsedHourCloseMarket.Hour(), parsedHourCloseMarket.Minute(),
					parsedHourCloseMarket.Second(), 0, now.Location())

				if fileModTime.After(marketCloseToday) {
					return
				} else {
					createBaseFile(fundsFile, baseFundsJson)
					createBaseFile(myFundsFile, baseMyFundsJson)
					updateValues()
				}
			}
		}
	default:
		log.Println("Too many arguments.")
	}
}
