package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func menu() {
	for {
		fmt.Println("Mutual funds options, use numbers only(1, 2, 3, etc)")
		fmt.Println("1- My funds")
		fmt.Println("2- All funds")
		fmt.Println("3- Exit")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch {
			case opt == "1":
				optionsMenu(myFundsFile)
			case opt == "2":
				optionsMenu(fundsFile)
			case opt == "3":
				innerFor := true
				for innerFor {
					fmt.Println("You sure want to exit? y/n")
					fmt.Print("> ")

					opt, err := reader.ReadString('\n')
					opt = strings.TrimSuffix(opt, "\n")
					if err == nil {
						switch opt {
						case "y", "Y":
							os.Exit(0)
						case "n", "N":
							innerFor = false
						default:
							fmt.Println("Wrong option")
						}
					} else {

						fmt.Printf("Error reading input: %v", err)
					}
				}

			default:
				fmt.Println("opt:", opt)
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func optionsMenu(file string) {
	for {
		fmt.Println("1- Show data")
		fmt.Println("2- Export data")
		fmt.Println("3- Modify data")
		fmt.Println("4- Back")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				menuShow(file)
			case "2":
				menuExport(file)
			case "3":
				menuModify(file)
			case "4":
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func menuShow(file string) {
	for {
		fmt.Println("1- Show data")
		fmt.Println("2- Back")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				showData(file)
			case "2":
				// TODO: once everything work erase this print
				fmt.Println("going back")
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func menuExport(file string) {
	for {
		fmt.Println("1- Export data")
		fmt.Println("2- Back")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				fmt.Println("exporting data")
			case "2":
				fmt.Println("going back")
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
	// TODO: Give options of exporting to a place on the pc choosen by the user
	// first option will be a .json file, later on maybe other options
}

func menuModify(file string) {
	for {
		fmt.Println("1- Modify data")
		fmt.Println("2- Back")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				fmt.Println("modifying data")
			case "2":
				fmt.Println("going back")
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}

	// TODO: This is difficult, ading a new fund should be easy, actually changing
	// or erasing a fund is more diffucult i think
}
