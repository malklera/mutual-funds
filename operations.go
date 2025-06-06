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

func showData(choosenFunds string) {
	data, err := os.ReadFile(fundsFile)
	if err != nil {
		log.Fatalf("Error reading file %s : %v", fundsFile, err)
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
	}

	var myFunds []Portfolio
	switch choosenFunds {
	case "myFunds":
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Fatalf("Error reading file %s : %v", myFundsFile, err)
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Fatalf("Error unmarshaling myFunds: %v", err)
		}
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
	case "all":
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
	default:
		// NOTE: Here i have the option of showing only one fund, on menuShow i
		// will have to add another option where you pass the name of the fund
		// and call showData with that name, for now i only will use that method
		// to modify, so the fund.Value, fund.Risk, and the calculation of lasValue
		// are not show here because the user cant change that
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Fatalf("Error reading file %s : %v", myFundsFile, err)
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Fatalf("Error unmarshaling myFunds: %v", err)
		}
		for _, fund := range funds {
			if fund.Name == choosenFunds {
				for _, myFund := range myFunds {
					if fund.Name == myFund.Name {
						fmt.Println("Name:", fund.Name)
						fmt.Println("Url:", fund.Url)
						fmt.Println("Owned shares:", myFund.Shares)
					}
				}
			}
		}

	}
}

func fundExist(fundName string) bool {
	data, err := os.ReadFile(fundsFile)
	if err != nil {
		log.Fatalf("Error reading file %s : %v", fundsFile, err)
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
	}
	
	for _, fund := range(funds) {
		if fundName == fund.Name {
			return true
		}
	}
	return false
}
