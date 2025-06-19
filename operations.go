package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Portfolio struct {
	Name   string  `json:"name"`
	Shares float64 `json:"shares"`
}

func showData(context string, choosenFunds string) error {
	data, err := os.ReadFile(fundsFile)
	if err != nil {
		log.Printf("Error reading file %s\n", fundsFile)
		return err
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		fmt.Print("Error unmarshaling data\n")
		return err
	}

	var myFunds []Portfolio
	switch context {
	case myFundsFile:
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", myFundsFile)
			return err
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Print("Error unmarshaling myFunds")
			return err
		}

		if choosenFunds == "allFunds" {
			for _, fund := range funds {
				for _, myFund := range myFunds {
					if fund.Name == myFund.Name {
						fmt.Println("Name:", fund.Name)
						fmt.Println("Url:", fund.URL)
						fmt.Println("Risk:", fund.Risk)
						fmt.Println("Owned shares:", myFund.Shares)
						lastValue := fund.Value[len(fund.Value)-1].Price * myFund.Shares
						fmt.Printf("Value owned shares: %f\n\n", lastValue)
						fmt.Printf("\tFrom:    %s\n", fund.Value[0].Date)
						fmt.Printf("\tTo:      %s\n", fund.Value[len(fund.Value)-1].Date)
						fmt.Printf("\tYield: $ %f\n", fund.Value[0].Price-fund.Value[len(fund.Value)-1].Price)
					}
				}
			}
		} else {
			exists, err := fundExist(myFundsFile, choosenFunds)
			if err != nil {
				log.Printf("Error checking the existence of %s\n", choosenFunds)
				return err
			} else {
				if exists {
					for _, fund := range funds {
						for _, myFund := range myFunds {
							if choosenFunds == myFund.Name {
								fmt.Println("Name:", fund.Name)
								fmt.Println("Url:", fund.URL)
								fmt.Println("Risk:", fund.Risk)
								fmt.Println("Owned shares:", myFund.Shares)
								lastValue := fund.Value[len(fund.Value)-1].Price * myFund.Shares
								fmt.Printf("Value owned shares: %f\n\n", lastValue)
								fmt.Printf("\tFrom:    %s\n", fund.Value[0].Date)
								fmt.Printf("\tTo:      %s\n", fund.Value[len(fund.Value)-1].Date)
								fmt.Printf("\tYield: $ %f\n", fund.Value[0].Price-fund.Value[len(fund.Value)-1].Price)
								return nil
							}
						}
					}
				} else {
					fmt.Printf("ChoosenFund: '%s'\n", choosenFunds)
					return errors.New("Fund do not exist on my funds")
				}
			}
		}
	case fundsFile:
		if choosenFunds == "allFunds" {
			for _, fund := range funds {
				fmt.Println("Name:", fund.Name)
				fmt.Println("Url:", fund.URL)
				fmt.Println("Risk:", fund.Risk)
			}
		} else {
			exists, err := fundExist(fundsFile, choosenFunds)
			if err != nil {
				log.Printf("Error checking the existence of %s\n", choosenFunds)
				return err
			} else {
				if exists {
					for _, fund := range funds {
						if choosenFunds == fund.Name {
							fmt.Println("Name:", fund.Name)
							fmt.Println("Url:", fund.URL)
							fmt.Println("Risk:", fund.Risk)
							return nil
						}
					}
				} else {
					fmt.Printf("ChoosenFund: '%s'\n", choosenFunds)
					return errors.New("Fund do not exist on funds")
				}
			}
		}
	default:
		return errors.New("wrong context")
	}

	return nil
}

func fundExist(context string, fundName string) (bool, error) {
	data, err := os.ReadFile(context)
	if err != nil {
		log.Printf("Error reading file %s\n", context)
		return false, err
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		fmt.Print("Error unmarshaling data\n")
		return false, err
	}

	for _, fund := range funds {
		if fundName == fund.Name {
			return true, nil
		}
	}
	return false, nil
}

func modifyData(context string, fundName string) error {
	var reader = bufio.NewReader(os.Stdin)
	switch context {
	case myFundsFile:
		var myFunds []Portfolio
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", myFundsFile)
			return err
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Print("Error unmarshaling myFunds")
			return err
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
					fmt.Println("1- Back")
					fmt.Println("Plese use the following format.")
					fmt.Println("Dot (.) for thousand separator (12.345)")
					fmt.Println("Coma (,) for decimal separator (12,3456)")
					fmt.Println("New ammount of owned shares:")
					fmt.Print("> ")

					newShares, err := reader.ReadString('\n')

					if err != nil {
						log.Printf("Error reading input: %v", err)
					} else {
						if newShares == "1" {
							return nil
						}
						newShares = strings.TrimSpace(newShares)

						if strings.Contains(newShares, ".") {
							newShares = strings.ReplaceAll(newShares, ".", "")
						}

						if strings.Contains(newShares, ",") {
							newShares = strings.ReplaceAll(newShares, ",", ".")
							parsedNewShares, err := strconv.ParseFloat(newShares, 64)
							if err != nil {
								log.Printf("Error parsing input: %s : %v", newShares, err)
							} else {
								newMyFund.Shares = parsedNewShares
								newMyFunds = append(newMyFunds, newMyFund)
								innerFor = false
							}
						} else {
							fmt.Printf("Input: '%s' is the wrong format", newShares)
						}
					}
				}
			} else {
				newMyFunds = append(newMyFunds, myFund)
			}
		}

		updateFunds, err := json.MarshalIndent(newMyFunds, "", "\t")
		if err != nil {
			log.Print("Error marshaling json from updateFunds\n")
			return err
		}

		err = os.WriteFile(myFundsFile, updateFunds, 0666)
		if err != nil {
			log.Printf("Error writing file %s\n", myFundsFile)
			return err
		}

	case fundsFile:
		var funds []Fund
		data, err := os.ReadFile(fundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", fundsFile)
			return err
		}
		err = json.Unmarshal(data, &funds)
		if err != nil {
			log.Print("Error unmarshaling funds")
			return err
		}

		var newFunds []Fund
		innerFor := true
		for _, fund := range funds {
			if fundName == fund.Name {
				var newFund Fund
				fmt.Println("Name:", fund.Name)
				fmt.Println("Url:", fund.URL)

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
						opt = strings.TrimSpace(opt)

						switch opt {
						case "1":
							fmt.Println("Enter the new name of the fund")
							name, err := reader.ReadString('\n')
							if err != nil {
								log.Printf("Error reading input: %v", err)
							} else {
								name = strings.TrimSpace(name)
								newFund.Name = name
								newFund.URL = fund.URL
								newFund.Risk = fund.Risk
								newFund.Value = append(newFund.Value, fund.Value...)

								newFunds = append(newFunds, newFund)
								innerFor = false
							}
						case "2":
							fmt.Println("Enter the new url of the fund")
							url, err := reader.ReadString('\n')
							if err != nil {
								log.Printf("Error reading input: %v", err)
							} else {
								url = strings.TrimSpace(url)
								newFund.Name = fund.Name
								newFund.URL = url
								newFund.Risk = fund.Risk
								newFund.Value = append(newFund.Value, fund.Value...)

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
			log.Print("Error marshaling json from updateFunds")
			return err
		}

		err = os.WriteFile(fundsFile, updateFunds, 0666)
		if err != nil {
			log.Printf("Error writing file %s\n", fundsFile)
			return err
		}

	default:
		return errors.New("wrong context")
	}
	return nil
}

func addData(context string, nameFund string) error {
	var reader = bufio.NewReader(os.Stdin)
	switch context {
	case fundsFile:
		fund := Fund{Name: nameFund}
		for {
			fmt.Println("DANGER: This data is not check")
			fmt.Println("Url:")
			fmt.Print("> ")

			url, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %s", err)
			} else {
				url = strings.TrimSpace(url)
				// NOTE: It seems a litle overboard to make a function to use once
				exist, resName, err := validURL(nameFund, url)
				if err != nil {
					fmt.Printf("Error running validURL: %v", err)
				} else {
					if exist {
						fund.URL = url
						break
					} else {
						fmt.Println("Discrepancy of names")
						fmt.Println("Given name:", nameFund)
						fmt.Println("Fund name from url:", resName)
					}
				}
			}
		}
		var funds []Fund
		data, err := os.ReadFile(fundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", fundsFile)
			return err
		}
		err = json.Unmarshal(data, &funds)
		if err != nil {
			log.Print("Error unmarshaling funds")
			return err
		}

		funds = append(funds, fund)
		updatedFunds, err := json.MarshalIndent(funds, "", "\t")
		if err != nil {
			log.Print("Error marshaling json from updatedFunds")
			return err
		}

		err = os.WriteFile(fundsFile, updatedFunds, 0666)
		if err != nil {
			log.Printf("Error writing file %s\n", fundsFile)
			return err
		}
	case myFundsFile:
		exists, err := fundExist(fundsFile, nameFund)
		if !exists {
			log.Printf("Error, fund: '%s' do not exist", nameFund)
			return err
		}
		myFund := Portfolio{Name: nameFund}
		for {
			fmt.Println("DANGER: This data is not check")
			fmt.Println("Plese use the following format.")
			fmt.Println("Dot (.) for thousand separator (12.345)")
			fmt.Println("Coma (,) for decimal separator (12,3456)")
			fmt.Println("Shares:")
			fmt.Print("> ")

			shares, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %s", err)
			} else {
				shares = strings.TrimSpace(shares)

				if strings.Contains(shares, ".") {
					shares = strings.ReplaceAll(shares, ".", "")
				}

				if strings.Contains(shares, ",") {
					shares = strings.ReplaceAll(shares, ",", ".")
					parsedShares, err := strconv.ParseFloat(shares, 64)
					if err != nil {
						fmt.Printf("Error parsing input: %s : %v", shares, err)
					} else {
						myFund.Shares = parsedShares
						break
					}
				} else {
					fmt.Printf("Input: '%s' is the wrong format", shares)
				}
			}
		}

		var myFunds []Portfolio
		myData, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", myFundsFile)
			return err
		}
		err = json.Unmarshal(myData, &myFunds)
		if err != nil {
			log.Print("Error unmarshaling myFunds")
			return err
		}

		myFunds = append(myFunds, myFund)

		updatedFunds, err := json.MarshalIndent(myFunds, "", "\t")
		if err != nil {
			log.Print("Error marshaling json from updatedFunds")
			return err
		}

		err = os.WriteFile(myFundsFile, updatedFunds, 0666)
		if err != nil {
			log.Printf("Error writing file %s\n", myFundsFile)
			return err
		}

	default:
		return errors.New("wrong context")
	}
	return nil
}

func deleteData(context string, nameFund string) error {
	var reader = bufio.NewReader(os.Stdin)
	confirmation := true
	for confirmation {
		fmt.Printf("DANGER: Are you sure you want to delete '%s' from '%s'?(y/n)", nameFund, context)
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %s", err)
		} else {
			opt = strings.TrimSpace(opt)
			switch opt {
			case "y", "Y":
				confirmation = false
			case "n", "N":
				return nil
			default:
				fmt.Println("Wrong option")
			}
		}
	}

	switch context {
	case fundsFile:
		var funds []Fund
		data, err := os.ReadFile(fundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", fundsFile)
			return err
		}
		err = json.Unmarshal(data, &funds)
		if err != nil {
			log.Print("Error unmarshaling funds\n")
			return err
		}

		var newFunds []Fund
		for _, fund := range funds {
			if fund.Name == nameFund {
				continue
			} else {
				newFunds = append(newFunds, fund)
			}
		}
		updatedFunds, err := json.MarshalIndent(newFunds, "", "\t")
		if err != nil {
			log.Print("Error marshaling json from updatedFunds")
			return err
		}

		err = os.WriteFile(fundsFile, updatedFunds, 0666)
		if err != nil {
			log.Printf("Error writing file %s\n", fundsFile)
			return err
		}
	case myFundsFile:
		var myFunds []Portfolio
		data, err := os.ReadFile(myFundsFile)
		if err != nil {
			log.Printf("Error reading file %s\n", myFundsFile)
			return err
		}
		err = json.Unmarshal(data, &myFunds)
		if err != nil {
			log.Print("Error unmarshaling funds\n")
			return err
		}

		var newMyFunds []Portfolio
		for _, myFund := range myFunds {
			if myFund.Name == nameFund {
				continue
			} else {
				newMyFunds = append(newMyFunds, myFund)
			}
		}
		updatedFunds, err := json.MarshalIndent(newMyFunds, "", "\t")
		if err != nil {
			log.Print("Error marshaling json from updatedFunds\n")
			return err
		}

		err = os.WriteFile(myFundsFile, updatedFunds, 0666)
		if err != nil {
			log.Printf("Error writing file %s\n", myFundsFile)
			return err
		}
	default:
		return errors.New("wrong context")
	}
	return nil
}

func exportData(context string, path string, choosenFunds string) error {
	if choosenFunds == "allFunds" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Print("Error getting current directory\n")
			return err
		}

		srcFile, err := os.Open(filepath.Join(currentDir, context))
		if err != nil {
			log.Printf("Error opening file %s\n", context)
			return err
		}
		defer func() {
			err := srcFile.Close()
			if err != nil {
				log.Printf("Error closing the file %s : %v", context, err)
			}
		}()

		destFile, err := os.Create(filepath.Join(path, context))
		if err != nil {
			log.Printf("Error creating file %s\n", context)
			return err
		}
		defer func() {
			err := destFile.Close()
			if err != nil {
				log.Printf("Error closing the file %s : %v", context, err)
			}
		}()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			log.Print("Error copying file\n")
			return err
		}
	} else {
		switch context {
		case fundsFile:
			var funds []Fund
			data, err := os.ReadFile(context)
			if err != nil {
				log.Printf("Error reading file %s\n", context)
				return err
			}
			err = json.Unmarshal(data, &funds)
			if err != nil {
				log.Print("Error unmarshaling data\n")
				return err
			}

			var exportFund Fund
			for _, fund := range funds {
				if fund.Name == choosenFunds {
					exportFund = fund
				}
			}

			marshaledFund, err := json.MarshalIndent(exportFund, "", "\t")
			if err != nil {
				log.Print("Error marshaling json from exportFund\n")
				return err
			}

			err = os.WriteFile(filepath.Join(path, context), marshaledFund, 0666)
			if err != nil {
				log.Printf("Error writing file %s\n", context)
				return err
			}
		case myFundsFile:
			// NOTE: think about this, i want to export only what is in myFunds.json
			// or export that and the related info from funds.json?
			// give option??
			var myFunds []Portfolio
			data, err := os.ReadFile(context)
			if err != nil {
				log.Printf("Error reading file %s\n", context)
				return err
			}
			err = json.Unmarshal(data, &myFunds)
			if err != nil {
				log.Print("Error unmarshaling data\n")
				return err
			}

			var exportFund Portfolio
			for _, myFund := range myFunds {
				if myFund.Name == choosenFunds {
					exportFund = myFund
				}
			}

			marshaledFund, err := json.MarshalIndent(exportFund, "", "\t")
			if err != nil {
				log.Print("Error marshaling json from exportFund\n")
				return err
			}

			err = os.WriteFile(filepath.Join(path, context), marshaledFund, 0666)
			if err != nil {
				log.Printf("Error writing file %s\n", context)
				return err
			}
		default:
			return errors.New("wrong context")
		}
	}
	return nil
}
