package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Portfolio struct {
	Name   string  `json:"name"`
	Shares float64 `json:"shares"`
}

func showData(context string, choosenFunds string) {
	data, err := os.ReadFile(fundsFile)
	if err != nil {
		// NOTE: i can deal with this error better by retrying or something
		log.Fatalf("Error reading file %s : %v", fundsFile, err)
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		// NOTE: i can deal with this error better by retrying or something
		fmt.Println("Error unmarshaling data:", err)
	}

	var myFunds []Portfolio
	switch context {
	case myFundsFile:
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Fatalf("Error reading file %s : %v", myFundsFile, err)
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Fatalf("Error unmarshaling myFunds: %v", err)
		}

		if choosenFunds == "allFunds" {
			for _, fund := range funds {
				for _, myFund := range myFunds {
					if fund.Name == myFund.Name {
						fmt.Println("Name:", fund.Name)
						fmt.Println("Url:", fund.Url)
						fmt.Println("Risk:", fund.Risk)
						fmt.Println("Value:")
						for _, value := range fund.Value {
							fmt.Println("\tDate:", value.Date)
							fmt.Printf("\tPrice: %.3f\n\n", value.Price)
						}
						fmt.Println("Owned shares:", myFund.Shares)
						lastValue := fund.Value[len(fund.Value)-1].Price * myFund.Shares
						fmt.Printf("Value shares: %f\n\n", lastValue)

						// WARNING: Is this correct? how is the yield calculated, call
						// the banc or read online, for now i will comment it

						//yield := fund.Value[0].Price - fund.Value[len(fund.Value)-1].Price
						//fmt.Println("Yield:", yield)
					}
				}
			}
		} else {
			for _, fund := range funds {
				for _, myFund := range myFunds {
					if choosenFunds == myFund.Name {
						fmt.Println("Name:", fund.Name)
						fmt.Println("Url:", fund.Url)
						fmt.Println("Risk:", fund.Risk)
						fmt.Println("Value:")
						for _, value := range fund.Value {
							fmt.Println("\tDate:", value.Date)
							fmt.Printf("\tPrice: %.3f\n\n", value.Price)
						}
						fmt.Println("Owned shares:", myFund.Shares)
						lastValue := fund.Value[len(fund.Value)-1].Price * myFund.Shares
						fmt.Printf("Value shares: %f\n\n", lastValue)

						return

						// WARNING: Is this correct? how is the yield calculated, call
						// the banc or read online, for now i will comment it

						//yield := fund.Value[0].Price - fund.Value[len(fund.Value)-1].Price
						//fmt.Println("Yield:", yield)
					}
				}
			}
		}
	case fundsFile:
		if choosenFunds == "allFunds" {
			for _, fund := range funds {
				fmt.Println("Name:", fund.Name)
				fmt.Println("Url:", fund.Url)
				fmt.Println("Risk:", fund.Risk)
				fmt.Println("Value:")
				for _, value := range fund.Value {
					fmt.Println("\tDate:", value.Date)
					fmt.Printf("\tPrice: %.3f\n\n", value.Price)
				}
			}
		} else {
			for _, fund := range funds {
				if choosenFunds == fund.Name {
					fmt.Println("Name:", fund.Name)
					fmt.Println("Url:", fund.Url)
					fmt.Println("Risk:", fund.Risk)
					fmt.Println("Value:")
					for _, value := range fund.Value {
						fmt.Println("\tDate:", value.Date)
						fmt.Printf("\tPrice: %.3f\n\n", value.Price)
					}
					return
				}
			}

		}
	default:
		log.Fatalf("Error, wrong context: %s", context)
	}
}

func fundExist(context string, fundName string) bool {
	data, err := os.ReadFile(context)
	if err != nil {
		// NOTE: i can deal with this error better by retrying or something
		log.Fatalf("Error reading file %s : %v", context, err)
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		// NOTE: i can deal with this error better by retrying or something
		fmt.Println("Error unmarshaling data:", err)
	}

	for _, fund := range(funds) {
		if fundName == fund.Name {
			return true
		}
	}
	return false
}
