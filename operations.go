package main

import (
	"bufio"
	"strings"
	"strconv"
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

	for _, fund := range funds {
		if fundName == fund.Name {
			return true
		}
	}
	return false
}

func modifyData(context string, fundName string) {
	// WARN: if i modified the fundsFile it will interfere with my checks of last
	// modified time, for know let it...
	var reader = bufio.NewReader(os.Stdin)
	switch context {
	case myFundsFile:
		var myFunds []Portfolio
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Fatalf("Error reading file %s : %v", myFundsFile, err)
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Fatalf("Error unmarshaling myFunds: %v", err)
		}

		var newMyFunds []Portfolio
		innerFor := true
		for _, myFund := range myFunds {
			if fundName == myFund.Name {
				var newMyFund Portfolio
				newMyFund.Name = myFund.Name
				fmt.Println("Name:", myFund.Name)
				fmt.Println("Owned shares:", myFund.Shares)
				for innerFor {
					fmt.Println("New ammount of owned shares:")
					fmt.Print("> ")

					newShares, err := reader.ReadString('\n')

					if err != nil {
						log.Printf("Error reading input: %v", err)
					} else {
						newShares = strings.TrimSuffix(newShares, "\n")
						parsedNewShares, err := strconv.ParseFloat(newShares, 64)
						if err != nil {
							log.Printf("Error parsing input: %v", err)
						} else {
							newMyFund.Shares = parsedNewShares
							newMyFunds = append(newMyFunds, newMyFund)
							innerFor = false
						}
					}
				}
			} else {
				newMyFunds = append(newMyFunds, myFund)
			}
		}

		updateFunds, err := json.MarshalIndent(newMyFunds, "", "\t")
		if err != nil {
			log.Fatalf("Error marshaling json from updateFunds: %v", err)
		}

		err = os.WriteFile(myFundsFile, updateFunds, 0666)
		if err != nil {
			log.Fatalf("Error writing file %s : %v", myFundsFile, err)
		}

	case fundsFile:
		var funds []Fund
		data, err := os.ReadFile(fundsFile)
		if err != nil {
			log.Fatalf("Error reading file %s : %v", fundsFile, err)
		}
		err = json.Unmarshal(data, &funds)
		if err != nil {
			log.Fatalf("Error unmarshaling funds: %v", err)
		}

		var newFunds []Fund
		innerFor := true
		for _, fund := range funds {
			if fundName == fund.Name {
				var newFund Fund
				fmt.Println("Name:", fund.Name)
				fmt.Println("Url:", fund.Url)

				for innerFor {
					fmt.Println("What do you which to change?")
					fmt.Println("DANGER: Neither is checked")
					fmt.Println("1- Name")
					fmt.Println("2- Url")
					fmt.Print("> ")

					opt, err := reader.ReadString('\n')

					if err != nil {
						log.Printf("Error reading input: %v", err)
					} else {
						opt = strings.TrimSuffix(opt, "\n")

						switch opt {
						case "1":
							fmt.Println("Enter the new name of the fund")
							name, err := reader.ReadString('\n')
							if err != nil {
								log.Printf("Error reading input: %v", err)
							} else {
								name = strings.TrimSuffix(name, "\n")
								newFund.Name = name
								newFund.Url = fund.Url
								newFund.Risk = fund.Risk
								for _, val := range fund.Value {
									newFund.Value = append(newFund.Value, val)
								}
								newFunds = append(newFunds, newFund)
								innerFor = false
							}
						case "2":
							fmt.Println("Enter the new url of the fund")
							url, err := reader.ReadString('\n')
							if err != nil {
								log.Printf("Error reading input: %v", err)
							} else {
								url = strings.TrimSuffix(url, "\n")
								newFund.Name = fund.Name
								newFund.Url = url
								newFund.Risk = fund.Risk
								for _, val := range fund.Value {
									newFund.Value = append(newFund.Value, val)
								}
								newFunds = append(newFunds, newFund)
								innerFor = false
							}
						default:
							fmt.Println("Wrong option")
						}
					}
				}
			} else {
				newFunds = append(newFunds, fund)
			}
		}
		updateFunds, err := json.MarshalIndent(newFunds, "", "\t")
		if err != nil {
			log.Fatalf("Error marshaling json from updateFunds: %v", err)
		}

		err = os.WriteFile(fundsFile, updateFunds, 0666)
		if err != nil {
			log.Fatalf("Error writing file %s : %v", fundsFile, err)
		}

	default:
		log.Fatalf("Error, wrong context: %s", context)
	}
}
